package main

import (
	"fmt"

	"github.com/Mrcrypt/5words/domain"
	"github.com/Mrcrypt/5words/mnemonic"
	"github.com/Mrcrypt/5words/msgswap"
	"github.com/Mrcrypt/5words/protoio"
	"github.com/c-bata/go-prompt"
	"github.com/libp2p/go-libp2p-core/peer"
	"google.golang.org/protobuf/proto"
)

var quit = prompt.KeyBind{
	Key: prompt.ControlC,
	Fn:  func(_ *prompt.Buffer) { panic("exit") },
}

func startCLI() {
	//State

	for {
		t := prompt.Input("> ", func(d prompt.Document) []prompt.Suggest {
			s := []prompt.Suggest{
				{Text: "addDomain", Description: "Adds a Domain to your list"},
				{Text: "msg", Description: "Send a message to a Domain from your list"},
				{Text: "quit", Description: "Quits"},
			}
			return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
		}, prompt.OptionAddKeyBind(quit), prompt.OptionShowCompletionAtStart())

		switch t {
		case "quit", "exit":
			return
		case "msg":
			t := prompt.Input("> Who do you want to message?: ", func(d prompt.Document) []prompt.Suggest {
				suggestions := []prompt.Suggest{}
				for _, contact := range store.contacts() {
					mnemonics, _ := contact.DomainEntry.Address.Mnemonics()
					suggestions = append(suggestions, prompt.Suggest{Text: contact.Alias, Description: mnemonics})
				}
				return prompt.FilterFuzzy(suggestions, d.GetWordBeforeCursor(), true)
			}, prompt.OptionShowCompletionAtStart())

			var mnemonic string
			for _, contact := range store.contacts() {
				if contact.Alias == t {
					mnemonic, _ = contact.DomainEntry.Address.Mnemonics()
					break
				}
			}
			if mnemonic == "" {
				fmt.Println("Alias not found")
				continue
			}

			dom, _ := domain.NameFromMnemonics(mnemonic)

			logger.Info("SEARCH KEY")
			ch, err := kademliaDHT.SearchValue(ctx, fmt.Sprintf("/peerdns/%016x", dom))
			if err != nil {
				panic(err)
			}
			logger.Info("SEARCHED KEY")

			res, ok := <-ch
			if !ok {
				logger.Info("Key not found")
				continue
			}
			logger.Info("FOUND KEY")

			var entry domain.Entry
			proto.Unmarshal(res, &entry)

			pid, err := peer.Decode(entry.Peers[0])
			if err != nil {
				panic(err)
			}

			stream, err := myHost.NewStream(ctx, pid, msgswap.ID)
			text := prompt.Input("> Write your Message:", func(_ prompt.Document) []prompt.Suggest { return nil })

			w := protoio.NewDelimitedWriter(stream)
			msg := msgswap.Message{
				Content: []byte(text),
				Topic:   []byte("Test Message"),
				TopicID: nil,
				Time:    nil,
				Address: &domain.Address{
					Nonce:  0,
					PubKey: nil,
				},
				Signature: nil,
			}
			w.WriteMsg(&msg)

		case "addDomain":
			words := prompt.Input("> Mnemonics: ", mnemonicsSuggestor)
			dom, _ := domain.NameFromMnemonics(words)

			logger.Info("SEARCH KEY")
			ch, err := kademliaDHT.SearchValue(ctx, fmt.Sprintf("/peerdns/%016x", dom))
			if err != nil {
				panic(err)
			}
			logger.Info("SEARCHED KEY")

			res, ok := <-ch
			if !ok {
				continue
			}
			logger.Info("FOUND KEY")

			alias := prompt.Input("> How do you name this Domain?:", func(_ prompt.Document) []prompt.Suggest { return nil })

			var entry domain.Entry
			proto.Unmarshal(res, &entry)

			store.saveContact(&Contact{Alias: alias, DomainEntry: &entry})
		}
	}
}

func mnemonicsSuggestor(d prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{}
	for _, word := range mnemonic.Words {
		suggestions = append(suggestions, prompt.Suggest{Text: word})
	}

	return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
}
