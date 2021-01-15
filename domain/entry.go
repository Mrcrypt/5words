package domain

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"

	ic "github.com/libp2p/go-libp2p-core/crypto"
	"google.golang.org/protobuf/proto"
)

//NewEntry ...
func NewEntry(privateAddress *PrivateAddress) (*Entry, error) {
	pubRSAKey, err := privateAddress.PublicRSAKey()
	if err != nil {
		return nil, err
	}

	return &Entry{
		Address:   &privateAddress.Address,
		PubRSAKey: pubRSAKey,
	}, nil
}

//Verify ...
func (d *Entry) Verify() (bool, error) {
	if err := d.Address.Verify(); err != nil {
		return false, err
	}
	pubKey, _ := ic.UnmarshalEd25519PublicKey(d.Address.PubKey)

	valid, _ := pubKey.Verify(d.signingData(), d.Signature)
	if !valid {
		return false, errors.New("Invalid Signature")
	}

	return true, nil
}

//signingData ...
func (d *Entry) signingData() []byte {
	var binaryNonce [8]byte
	binary.BigEndian.PutUint64(binaryNonce[:], d.Address.Nonce)

	valuesBytes, _ := json.Marshal(d.Values)
	peersBytes, _ := json.Marshal(d.Peers)

	return bytes.Join([][]byte{
		binaryNonce[:],
		d.Address.PubKey,
		valuesBytes,
		peersBytes,
	}, []byte{})
}

//Sign ...
func (d *Entry) Sign(privateAddress *PrivateAddress) error {
	privKey, err := ic.UnmarshalPrivateKey(privateAddress.PrivateKey)
	if err != nil {
		return err
	}

	sign, err := privKey.Sign(d.signingData())
	if err != nil {
		return err
	}
	d.Signature = sign
	return nil
}

//UnmarshalYAML ...
func (d *Entry) UnmarshalYAML(unmarshal func(interface{}) error) error {
	data := []byte{}
	unmarshal(&data)
	return proto.Unmarshal(data, d)
}

//MarshalYAML ...
func (d *Entry) MarshalYAML() (interface{}, error) {
	return proto.Marshal(d)
}
