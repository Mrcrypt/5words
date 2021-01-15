package main

import (
	"errors"
	"fmt"

	"github.com/Mrcrypt/5words/domain"
	record "github.com/libp2p/go-libp2p-record"
	"google.golang.org/protobuf/proto"
)

var _ record.Validator = &PeerDNSValidator{}

//PeerDNSValidator ...
type PeerDNSValidator struct {
}

// Validate validates the given record, returning an error if it's
// invalid (e.g., expired, signed by the wrong key, etc.).
func (p *PeerDNSValidator) Validate(key string, value []byte) error {
	var dom domain.Entry
	proto.Unmarshal(value, &dom)

	_, err := dom.Verify()
	if err != nil {
		return err
	}

	name, err := dom.Address.Name()
	if err != nil {
		return err
	}
	if key != fmt.Sprintf("/peerdns/%016x", name) {
		return errors.New("Wrong Key")
	}

	return err
}

// Select selects the best record from the set of records (e.g., the
// newest).
//
// Decisions made by select should be stable.
func (p *PeerDNSValidator) Select(key string, values [][]byte) (int, error) {
	return 0, nil
}
