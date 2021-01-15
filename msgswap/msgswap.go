package msgswap

//go:generate protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I. -I.. message.proto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/binary"
	"errors"
	"fmt"

	domain "github.com/Mrcrypt/5words/domain"
	"github.com/Mrcrypt/5words/protoio"
	ic "github.com/libp2p/go-libp2p-core/crypto"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/protocol"
	"google.golang.org/protobuf/proto"
)

//ID ...
var ID = protocol.ID("/msgSwap/1.0.0")

//HandleStream ...
func HandleStream(callback func(*EncryptedMessage)) network.StreamHandler {
	return func(stream network.Stream) {
		fmt.Println("Got a new stream!")

		reader := protoio.NewDelimitedReader(stream, 80*1024)

		m := EncryptedMessage{}
		reader.ReadMsg(&m)

		fmt.Printf("You Recieved a Message\n")
		callback(&m)

	}
}

//StoredMessage ...
type StoredMessage struct {
	ID string
	Message
}

//NewMessage ...
func NewMessage(from ic.PrivKey, nonce uint64, to string, content string) {

}

//Encrypt ...
func (m *Message) Encrypt(pubRSAKey []byte) (*EncryptedMessage, error) {
	pub, err := x509.ParsePKIXPublicKey(pubRSAKey)
	if err != nil {
		return nil, err
	}
	pk, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not actually an rsa public key")
	}

	data, err := proto.Marshal(m)
	if err != nil {
		return nil, err
	}

	//AES Encyption
	key := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(key); err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	aesGCM, err := cipher.NewGCM(block)
	nonce := make([]byte, aesGCM.NonceSize())
	ciphertext := aesGCM.Seal(nonce, nonce, data, nil)

	//RSA Encryption
	encryptedKey, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, pk, key, nil)
	if err != nil {
		return nil, err
	}

	return &EncryptedMessage{
		Reciever:     m.Reciever,
		EncryptedKey: encryptedKey,
		Message:      ciphertext,
	}, nil
}

//Decrypt ...
func (m *EncryptedMessage) Decrypt(privRSAKey []byte) (*Message, error) {
	priv, err := x509.ParsePKCS1PrivateKey(privRSAKey)
	if err != nil {
		return nil, err
	}

	//RSA Decryption
	decryptedKey, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, priv, m.EncryptedKey, nil)
	if err != nil {
		return nil, err
	}

	//AES Decryption
	block, err := aes.NewCipher(decryptedKey)
	aesGCM, err := cipher.NewGCM(block)

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := m.Message[:nonceSize], m.Message[nonceSize:]

	plaintext, err := aesGCM.Open(ciphertext[:0], nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	var msg Message
	err = proto.Unmarshal(plaintext, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

//Verify ...
func (m *Message) Verify() (bool, error) {
	if err := m.Address.Verify(); err != nil {
		return false, err
	}
	pubKey, _ := ic.UnmarshalEd25519PublicKey(m.Address.PubKey)

	valid, _ := pubKey.Verify(m.signingData(), m.Signature)
	if !valid {
		return false, errors.New("Invalid Signature")
	}

	return true, nil
}

//signingData ...
func (m *Message) signingData() []byte {
	var binaryNonce [8]byte
	binary.BigEndian.PutUint64(binaryNonce[:], m.Address.Nonce)

	return bytes.Join([][]byte{
		m.Content,
		m.Topic,
		m.Time,
		m.Reciever,
		m.Address.PubKey,
		binaryNonce[:],
	}, []byte{})
}

//Sign ...
func (m *Message) Sign(privateAddress *domain.PrivateAddress) error {
	privKey, err := ic.UnmarshalPrivateKey(privateAddress.PrivateKey)
	if err != nil {
		return err
	}

	sign, err := privKey.Sign(m.signingData())
	if err != nil {
		return err
	}
	m.Signature = sign
	return nil
}
