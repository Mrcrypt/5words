package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime/quotedprintable"
	"net/http"
	"time"

	"github.com/Mrcrypt/5words/domain"
	"github.com/Mrcrypt/5words/mnemonic"
	"github.com/Mrcrypt/5words/msgswap"
	"github.com/Mrcrypt/5words/protoio"
	"github.com/julienschmidt/httprouter"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/markbates/pkger"
	"google.golang.org/protobuf/proto"
)

func startWebAPI() {
	r := httprouter.New()

	r.GlobalOPTIONS = http.HandlerFunc(allowCORS)

	r.NotFound = http.FileServer(pkger.Dir("/webnode/dist"))

	r.GET("/wordlist", wordList)

	r.GET("/messages", messages)
	r.POST("/sendmessage", sendmessage)

	//Contacts
	r.GET("/contacts", contacts)
	r.POST("/newcontact", newContact)
	r.DELETE("/contacts", deleteContact)
	r.PUT("/contacts", updateContact)

	//Following
	r.POST("/newfollowing", newFollowing)
	r.GET("/following", following)
	r.GET("/followers", followers)

	//Accounts
	r.POST("/accounts", createAccount)
	r.DELETE("/accounts", deleteAccount)
	r.PUT("/accounts", updateAccount)
	r.GET("/accounts", getAccounts)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", config.APIPort),
		Handler: PreMiddleware{r},
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

//PreMiddleware ...
type PreMiddleware struct {
	Handler http.Handler
}

func (m PreMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	m.Handler.ServeHTTP(w, r)
}

func allowCORS(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Access-Control-Request-Method") != "" {
		// Set CORS headers
		header := w.Header()
		header.Set("Allow", "POST, GET, OPTIONS, DELETE, PUT")
		header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
		header.Set("Access-Control-Allow-Origin", "*")
		header.Set("Access-Control-Allow-Headers", "*")
	}

	// Adjust status code to 204
	w.WriteHeader(http.StatusNoContent)
}

func newContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jreq := struct {
		Words string
		Alias string
	}{}
	json.NewDecoder(r.Body).Decode(&jreq)

	dom, _ := domain.NameFromMnemonics(jreq.Words)

	logger.Info("SEARCH KEY")
	ch, err := kademliaDHT.SearchValue(ctx, fmt.Sprintf("/peerdns/%016x", dom))
	if err != nil {
		panic(err)
	}
	logger.Info("SEARCHED KEY")

	res, ok := <-ch
	if !ok {
		return
	}
	logger.Info("FOUND KEY")

	var entry domain.Entry
	proto.Unmarshal(res, &entry)

	store.saveContact(&Contact{Alias: jreq.Alias, DomainEntry: &entry})
}

func contacts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contacts := []interface{}{}
	for _, c := range store.contacts() {
		mnemonics, _ := c.DomainEntry.Address.Mnemonics()
		contacts = append(contacts, map[string]interface{}{
			"alias": c.Alias,
			"words": mnemonics,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"contacts": contacts,
	})
}

func deleteContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jreq := struct {
		Words string
	}{}
	json.NewDecoder(r.Body).Decode(&jreq)

	store.deleteContact(jreq.Words)
}

func updateContact(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jreq := struct {
		Words string
		Alias string
	}{}

	json.NewDecoder(r.Body).Decode(&jreq)

	contact := store.getContact(jreq.Words)
	if contact == nil {
		fmt.Println("Contact not found")
		return
	}
	contact.Alias = jreq.Alias
	store.saveContact(contact)
}

func messages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	messages := []interface{}{}
	for i, m := range store.messages() {
		b, _ := ioutil.ReadAll(quotedprintable.NewReader(bytes.NewReader(m.Content)))

		var msgTime time.Time
		msgTime.UnmarshalText(m.Time)

		mnemonics, _ := m.Address.Mnemonics()
		sender := mnemonics
		contact := store.getContact(mnemonics)
		if contact != nil {
			sender = contact.Alias
		}

		messages = append(messages, map[string]interface{}{
			"id":      i, //string(m.TopicID),
			"topic":   string(m.Topic),
			"name":    string(sender),
			"content": string(b),
			"time":    string(msgTime.Format("Mon, 2 Jan 06 15:04")),
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"messages": messages,
	})
}

func wordList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"wordList": mnemonic.Words,
	})
}

func sendmessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jreq := struct {
		Reciever []string
		Content  string
		Title    string
	}{}
	json.NewDecoder(r.Body).Decode(&jreq)

	fmt.Printf("reciever = %+v\n", jreq.Reciever)
	fmt.Printf("title = %+v\n", jreq.Title)
	fmt.Printf("content = %+v\n", jreq.Content)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"errors": nil,
	})

	sendToReceiver := func(mnemonic string) {
		fmt.Printf("mnemonic = %+v\n", mnemonic)
		dom, err := domain.NameFromMnemonics(mnemonic)
		if err != nil {
			fmt.Printf("message: %s", err)
			return
		}

		logger.Info("SEARCH KEY")
		ch, err := kademliaDHT.SearchValue(ctx, fmt.Sprintf("/peerdns/%016x", dom))
		if err != nil {
			return
		}
		logger.Info("SEARCHED KEY")

		res, ok := <-ch
		if !ok {
			logger.Info("Key not found")
			return
		}
		logger.Info("FOUND KEY")

		var entry domain.Entry
		proto.Unmarshal(res, &entry)

		pid, err := peer.Decode(entry.Peers[0])
		if err != nil {
			return
		}

		stream, err := myHost.NewStream(ctx, pid, msgswap.ID)
		if err != nil {
			fmt.Printf("err = %+v\n", err)
			return
		}
		logger.Info("Opened Stream")

		now, _ := time.Now().MarshalText()
		p2pw := protoio.NewDelimitedWriter(stream)
		msg := msgswap.Message{
			Content:   []byte(jreq.Content),
			Topic:     []byte(jreq.Title),
			Time:      now,
			Reciever:  dom,
			Address:   &config.PrivAddresses[0].Address,
			Signature: nil,
		}
		msg.Sign(&config.PrivAddresses[0])

		emsg, err := msg.Encrypt(entry.PubRSAKey)
		if err != nil {
			return
		}

		p2pw.WriteMsg(emsg)
		logger.Info("Wrote Message")

	}

	for _, reciever := range jreq.Reciever {

		var mnemonic string

		if reciever == "Follower" {
			followers := store.followers()

			for _, follower := range followers {
				fmt.Printf("follower = %+v\n", follower)
				fmt.Printf("follower.DomainEntry = %+v\n", follower.DomainEntry)
				mnemonics, _ := follower.DomainEntry.Address.Mnemonics()
				sendToReceiver(mnemonics)
			}
		} else {
			for _, contact := range store.contacts() {
				if contact.Alias == reciever {
					mnemonic, _ = contact.DomainEntry.Address.Mnemonics()
					break
				}
			}
			if mnemonic == "" {
				fmt.Println("Alias not found")
				return
			}
			sendToReceiver(mnemonic)
		}

	}
}

func newFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jreq := struct {
		Words string
		Alias string
	}{}
	json.NewDecoder(r.Body).Decode(&jreq)

	dom, _ := domain.NameFromMnemonics(jreq.Words)

	logger.Info("SEARCH KEY")
	ch, err := kademliaDHT.SearchValue(ctx, fmt.Sprintf("/peerdns/%016x", dom))
	if err != nil {
		panic(err)
	}
	logger.Info("SEARCHED KEY")

	res, ok := <-ch
	if !ok {
		return
	}
	logger.Info("FOUND KEY")

	var entry domain.Entry
	proto.Unmarshal(res, &entry)

	contact := Contact{Alias: jreq.Alias, DomainEntry: &entry}
	store.createFollowing(&contact)

	pid, err := peer.Decode(entry.Peers[0])
	if err != nil {
		panic(err)
	}

	stream, err := myHost.NewStream(ctx, pid, msgswap.FollowID)

	p2pw := protoio.NewDelimitedWriter(stream)

	name, _ := config.PrivAddresses[0].Name()
	msg := msgswap.NewFollower{
		Reciever:  nil,
		Name:      name,
		Nonce:     nil,
		PubKey:    nil,
		Signature: nil,
	}
	p2pw.WriteMsg(&msg)
}

func following(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contacts := store.following()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"following": contacts,
	})
}

func followers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	contacts := store.followers()

	json.NewEncoder(w).Encode(map[string]interface{}{
		"followers": contacts,
	})
}

func deleteAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	j := struct {
		Alias   string
		Address string
	}{}
	json.NewDecoder(r.Body).Decode(&j)

	for i, account := range config.PrivAddresses {
		mnemonics, _ := account.Mnemonics()
		if j.Address == mnemonics {
			config.PrivAddresses = append(config.PrivAddresses[:i], config.PrivAddresses[i+1:]...)
			saveConfig()
		}
	}

}

func updateAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	j := struct {
		Alias   string
		Address string
	}{}
	json.NewDecoder(r.Body).Decode(&j)

	for i, account := range config.PrivAddresses {
		mnemonics, _ := account.Mnemonics()
		if j.Address == mnemonics {
			config.PrivAddresses[i].Alias = j.Alias
			saveConfig()
		}
	}

}
func createAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	j := struct {
		Alias string
	}{}
	json.NewDecoder(r.Body).Decode(&j)

	dom, _ := domain.NewPrivateAddress(j.Alias)

	if config.PrivAddresses == nil {
		config.PrivAddresses = []domain.PrivateAddress{}
	}
	config.PrivAddresses = append(config.PrivAddresses, *dom)

	saveConfig()
}

func getAccounts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	accounts := []interface{}{}
	for _, account := range config.PrivAddresses {
		mnemonics, _ := account.Address.Mnemonics()
		accounts = append(accounts, map[string]interface{}{
			"alias":   account.Alias,
			"address": mnemonics,
		})
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"accounts": accounts,
	})
}
