package main

import (
	"os"

	"github.com/Mrcrypt/5words/domain"
	"gopkg.in/yaml.v2"
)

//Config ...
type Config struct {
	BootstrapPeers []string
	StorePath      string
	Port           uint16
	APIPort        uint16
	PrivAddresses  []domain.PrivateAddress `yaml:",flow"`
	PeerIDKey      []byte                  `yaml:",flow"`
	Relay          bool
}

//Contact ...
type Contact struct {
	Alias       string
	DomainEntry *domain.Entry
}

var configPath string
var config Config

func loadConfig() error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return err
	}
	return nil
}

func saveConfig() error {
	f, err := os.OpenFile(configPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := yaml.NewEncoder(f)
	err = encoder.Encode(&config)
	if err != nil {
		return err
	}
	return nil
}
