package main

import (
	"bytes"
	"context"
	"fmt"
	"net"
	sync "sync"
	"time"

	"github.com/Mrcrypt/5words/domain"
	"github.com/Mrcrypt/5words/msgswap"
	"github.com/google/uuid"
	dssync "github.com/ipfs/go-datastore/sync"
	"github.com/libp2p/go-libp2p"
	circuit "github.com/libp2p/go-libp2p-circuit"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/libp2p/go-libp2p-core/protocol"
	crypto "github.com/libp2p/go-libp2p-crypto"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	routing "github.com/libp2p/go-libp2p-routing"
	"github.com/multiformats/go-multiaddr"
)

var myHost host.Host
var kademliaDHT *dht.IpfsDHT
var datastore *PersistentDatastore
var ctx context.Context

func startNode() {
	var err error

	port := config.Port
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		port = 0
		fmt.Printf("Port %d in use, using random port.\n", config.Port)
	} else {
		lis.Close()
	}

	var privPeerIDKey crypto.PrivKey
	if config.PeerIDKey == nil || len(config.PeerIDKey) == 0 {
		privPeerIDKey, _, _ = crypto.GenerateKeyPair(crypto.Ed25519, 0)
		config.PeerIDKey, _ = crypto.MarshalPrivateKey(privPeerIDKey)
		saveConfig()
	} else {
		privPeerIDKey, _ = crypto.UnmarshalPrivateKey(config.PeerIDKey)
	}

	datastore = NewPersistentDatastore()

	var opthop libp2p.Option
	if config.Relay {
		opthop = libp2p.EnableRelay(circuit.OptHop)
	}

	myHost, err = libp2p.New(ctx,
		libp2p.ListenAddrStrings(
			fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port),
			fmt.Sprintf("/ip6/::/tcp/%d", port),
		),
		libp2p.Identity(privPeerIDKey),
		libp2p.NATPortMap(),
		libp2p.DefaultStaticRelays(),
		libp2p.EnableNATService(),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			var err error
			kademliaDHT, err = dht.New(ctx, h,
				dht.Validator(&PeerDNSValidator{}),
				//dht.Mode(dht.ModeServer),
				dht.ProtocolPrefix(protocol.ID("5words")),
				dht.Datastore(dssync.MutexWrap(datastore)),
				dht.MaxRecordAge(time.Hour*24*7),
			)
			return kademliaDHT, err
		}),
		libp2p.EnableAutoRelay(),
		opthop,
	)
	if err != nil {
		panic(err)
	}

	myHost.SetStreamHandler(msgswap.ID, msgswap.HandleStream(recievedMessage))
	myHost.SetStreamHandler(msgswap.FollowID, msgswap.FollowHandleStream(recievedFollower))

	for i := range myHost.Addrs() {
		logger.Info("We Are: ", myHost.Addrs()[i], "/p2p/", myHost.ID())
	}

	var wg sync.WaitGroup
	for _, ma := range config.BootstrapPeers {
		peerAddr, _ := multiaddr.NewMultiaddr(ma)
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := myHost.Connect(ctx, *peerinfo); err != nil {
				logger.Warning(err)
			} else {
				logger.Info("Connection established with bootstrap node:", *peerinfo)
			}
		}()
	}
	wg.Wait()

	logger.Debug("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}
}

func recievedMessage(emsg *msgswap.EncryptedMessage) {
	//Decrypt
	var privAddr *domain.PrivateAddress
	for i, addr := range config.PrivAddresses {
		name, _ := addr.Address.Name()
		if bytes.Compare(name, emsg.Reciever) == 0 {
			privAddr = &config.PrivAddresses[i]
		}
	}
	fmt.Printf("privAddr.Alias = %+v\n", privAddr.Alias)

	msg, err := emsg.Decrypt(privAddr.PrivateRSAKey)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
		return
	}

	valid, _ := msg.Verify()
	if valid == false {
		fmt.Println("Received invalid message")
		return
	}

	//Is Contact
	mnemonics, _ := msg.Address.Mnemonics()
	contact := store.getContact(mnemonics)
	if contact == nil {
		fmt.Println("Sender not in Contacts")
		return
	}

	id, _ := uuid.NewRandom()
	store.saveMessage(&msgswap.StoredMessage{ID: id.String(), Message: *msg})
}

func recievedFollower(newFollower *msgswap.NewFollower) {
	store.createFollower(&Contact{
		Alias: string(newFollower.Name),
		DomainEntry: &domain.Entry{
			Address: &domain.Address{
				Nonce:  0,
				PubKey: nil,
			},
			Signature: nil,
			Values: map[string]string{
				"": "",
			},
			Peers: nil,
		},
	})
}
