package main

import (
	"fmt"
	"testing"

	"github.com/dgraph-io/badger/v2"
)

func TestGetAllPrefixed(t *testing.T) {
	db, err := badger.Open(badger.DefaultOptions("").WithInMemory(true))
	if err != nil {
		panic(err)
	}

	s := &Store{
		db: db,
	}

	s.saveStruct("test/heinz", Contact{Alias: "asdf"})
	s.saveStruct("test/peter", Contact{Alias: "peter"})

	var contacts []Contact = []Contact{}
	s.getAllPrefixed("test", &contacts)
	fmt.Printf("contacts = %+v\n", contacts)
}
