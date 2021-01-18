package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/Mrcrypt/5words/domain"
	"github.com/ipfs/go-log/v2"
	"github.com/webview/webview"
	"google.golang.org/protobuf/proto"
)

var logger = log.Logger("5Words")

var store *Store

func main() {
	log.SetLogLevel("5Words", "info")

	//Flags
	configPathF := flag.String("config", "config.yml", "Config Path")
	debug := flag.Bool("debug", false, "")
	table := flag.Bool("table", false, "")
	gui := flag.Bool("gui", false, "Opens the GUI")
	cli := flag.Bool("cli", false, "Opens the CLI")
	bootstrap := flag.String("boot", "", "BootstrapPeer")
	flag.Parse()

	configPath = *configPathF

	if *debug {
		log.SetAllLoggers(log.LevelDebug)
	}

	//Node
	ctx = context.Background()

	loadConfig()
	//Database
	storePath := "./db"
	if config.StorePath != "" {
		storePath = config.StorePath
	}
	store = newStore(storePath)
	defer store.db.Close()

	//ONLY FOR TESTING
	if bootstrap != nil && *bootstrap != "" {
		var exists bool
		for _, p := range config.BootstrapPeers {
			if p == *bootstrap {
				exists = true
			}
		}
		if !exists {
			config.BootstrapPeers = append(config.BootstrapPeers, *bootstrap)
		}
	}

	saveConfig()

	//////////

	startNode()

	//TESTING ONLY
	if len(config.BootstrapPeers) == 0 {
		config.BootstrapPeers = []string{myHost.Addrs()[0].String() + "/p2p/" + myHost.ID().String()}
	}
	saveConfig()
	//////////

	//ONLY FOR TESTING
	if *table {
		go func() {
			for {
				//kademliaDHT.RoutingTable().Print()
				//datastore.PrintValues()
				for i := range myHost.Addrs() {
					logger.Info("We Are: ", myHost.Addrs()[i], "/p2p/", myHost.ID())
				}
				time.Sleep(10 * time.Second)
			}
		}()
	}
	//////////

	//Create Default Address
	if len(config.PrivAddresses) == 0 {
		addr, err := domain.NewPrivateAddress("Default")
		if err != nil {
			panic(err)
		}

		config.PrivAddresses = []domain.PrivateAddress{
			*addr,
		}
		saveConfig()
	}
	//ONLY FOR TESTING
	go func() {
		time.Sleep(1 * time.Second)
		for {
			for _, privAddr := range config.PrivAddresses {
				entry, _ := domain.NewEntry(&privAddr)
				entry.Peers = []string{myHost.ID().String()}

				mnemonics, _ := entry.Address.Mnemonics()
				logger.Infof("Our Domain: %s", mnemonics)

				entry.Sign(&privAddr)

				logger.Info("PUT VALUE")
				data, _ := proto.Marshal(entry)

				name, _ := entry.Address.Name()
				err := kademliaDHT.PutValue(ctx, fmt.Sprintf("/peerdns/%016x", name), data)
				if err != nil {
					fmt.Println(err)
				}
				logger.Info("PUTTED VALUE")
				time.Sleep(10 * time.Minute)
			}
		}
	}()
	//////////

	if config.APIPort != 0 {
		go startWebAPI()
	}
	if *cli == true {
		go startCLI()
	}
	if *gui == true {
		go startWebView()
	}
	select {}
}

func startWebView() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("5Words.io")
	w.SetSize(800, 600, webview.HintNone)
	w.Init(fmt.Sprintf(`var APIPort = %d;`, config.APIPort))
	w.Navigate(fmt.Sprintf("http://127.0.0.1:%d", config.APIPort))
	w.Run()
}
