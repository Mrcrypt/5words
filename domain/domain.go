package domain

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative domain.proto

import (
	"crypto/rand"
	"crypto/x509"
	"encoding/binary"
	"errors"
	"math/big"
	"strings"

	"github.com/Mrcrypt/5words/mnemonic"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"golang.org/x/crypto/argon2"
)

var difficulty = big.NewInt(100)
var target = new(big.Int).Div(new(big.Int).SetUint64(0xFFFFFFFFFFFFFFFF), difficulty)

//PowHash ...
func PowHash(data []byte) []byte {
	return argon2.Key(data, []byte("domain"), 1, 1024, 1, 16)
}

//PrivateAddress ...
type PrivateAddress struct {
	Address
	Alias         string
	PrivateKey    []byte
	PrivateRSAKey []byte
}

//NewPrivateAddress ...
func NewPrivateAddress(alias string) (*PrivateAddress, error) {
	privKey, _, err := ic.GenerateKeyPair(ic.Ed25519, 0)
	pubKey, err := privKey.GetPublic().Raw()
	if err != nil {
		return nil, err
	}
	privRSAKey, _, err := ic.GenerateKeyPair(ic.RSA, 2048)
	privRSAKeyBytes, err := privRSAKey.Raw()

	nonceBytes := make([]byte, 8)
	_, err = rand.Read(nonceBytes)
	if err != nil {
		return nil, err
	}

	nonce := binary.BigEndian.Uint64(nonceBytes)

	data := make([]byte, len(pubKey)+8)
	copy(data, pubKey)

	var hash []byte

	for {
		binary.BigEndian.PutUint64(data[len(pubKey):], nonce)
		hash = PowHash(data)

		result := new(big.Int).SetBytes(hash[:8])
		if result.Cmp(target) == -1 {
			break
		}

		nonce++
	}

	privKeyBytes, err := ic.MarshalPrivateKey(privKey)
	if err != nil {
		return nil, err
	}

	addr := &PrivateAddress{
		Address: Address{
			Nonce:  nonce,
			PubKey: pubKey,
		},
		Alias:         alias,
		PrivateKey:    privKeyBytes,
		PrivateRSAKey: privRSAKeyBytes,
	}

	return addr, nil
}

//PublicRSAKey ...
func (a *PrivateAddress) PublicRSAKey() ([]byte, error) {
	priv, err := x509.ParsePKCS1PrivateKey(a.PrivateRSAKey)
	if err != nil {
		return nil, err
	}

	return x509.MarshalPKIXPublicKey(&priv.PublicKey)
}

//Name of a Address in Bytes
func (a *Address) Name() ([]byte, error) {
	data := make([]byte, len(a.PubKey)+8)
	copy(data, a.PubKey)
	var hash []byte

	binary.BigEndian.PutUint64(data[len(a.PubKey):], a.Nonce)
	hash = PowHash(data)

	result := new(big.Int).SetBytes(hash[:8])
	if result.Cmp(target) != -1 {
		return nil, errors.New("Invalid Nonce")
	}

	return hash[8:], nil
}

//Verify a Address
func (a *Address) Verify() error {
	data := make([]byte, len(a.PubKey)+8)
	copy(data, a.PubKey)
	var hash []byte

	binary.BigEndian.PutUint64(data[len(a.PubKey):], a.Nonce)
	hash = PowHash(data)

	result := new(big.Int).SetBytes(hash[:8])
	if result.Cmp(target) != -1 {
		return errors.New("Invalid Nonce")
	}

	return nil
}

//Mnemonics of the Address
func (a *Address) Mnemonics() (string, error) {
	var result []string = make([]string, 0)

	wordListLen := big.NewInt(int64(len(mnemonic.Words)))

	nameBytes, err := a.Name()
	if err != nil {
		return "", err
	}

	name := new(big.Int).SetBytes(nameBytes)
	for i := 4; i >= 0; i-- {
		a := new(big.Int).Quo(name, new(big.Int).Exp(wordListLen, big.NewInt(int64(i)), nil))
		a.Mod(a, big.NewInt(int64(len(mnemonic.Words))))
		result = append(result, mnemonic.Words[a.Int64()])
	}

	return strings.Join(result, " "), nil
}

//NameFromMnemonics ...
func NameFromMnemonics(mnemonics string) ([]byte, error) {
	words := strings.Split(mnemonics, " ")
	if len(words) != 5 {
		return nil, errors.New("Invalid word count")
	}

	wordListLen := big.NewInt(int64(len(mnemonic.Words)))
	name := new(big.Int)
	for i := 0; i < 5; i++ {
		idx := 0
		for index, word := range mnemonic.Words {
			if word == words[i] {
				idx = index
				break
			}
		}

		offset := new(big.Int).Exp(wordListLen, big.NewInt(int64(4-i)), nil)
		name.Add(name, new(big.Int).Mul(offset, big.NewInt(int64(idx))))
	}

	return name.Bytes(), nil
}
