package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"reflect"

	"github.com/Mrcrypt/5words/msgswap"
	"github.com/dgraph-io/badger/v2"
)

//Store ...
type Store struct {
	db *badger.DB
}

func newStore(path string) *Store {
	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		panic(err)
	}

	return &Store{
		db: db,
	}
}

func (s *Store) saveStruct(key string, val interface{}) {
	s.db.Update(func(txn *badger.Txn) error {
		buf := bytes.Buffer{}
		enc := gob.NewEncoder(&buf)
		enc.Encode(val)
		txn.Set([]byte(key), buf.Bytes())
		return nil
	})
}

func (s *Store) getStruct(key string, val interface{}) error {
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(data []byte) error {
			reader := bytes.NewReader(data)
			dec := gob.NewDecoder(reader)
			err := dec.Decode(val)
			return err
		})
		return nil
	})
	return err
}

func (s *Store) getAllPrefixed(prefix string, output interface{}) error {
	sliceType := reflect.TypeOf(output)

	if sliceType.Kind() != reflect.Ptr {
		return errors.New("output must be a pointer to a slice")
	}

	sliceType = sliceType.Elem()

	if sliceType.Kind() != reflect.Slice {
		return errors.New("output must be a pointer to a slice")
	}
	sliceType = sliceType.Elem()

	s.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		pre := []byte(prefix)
		for it.Seek(pre); it.ValidForPrefix(pre); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				t := reflect.New(sliceType)
				gob.NewDecoder(bytes.NewReader(v)).Decode(t.Interface())
				sliceVal := reflect.ValueOf(output)
				newSlice := reflect.Append(reflect.Indirect(sliceVal), reflect.Indirect(t))
				sliceVal.Elem().Set(newSlice)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

func (s *Store) createFollowing(contact *Contact) {
	s.saveStruct("following/"+contact.Alias, contact)
}

func (s *Store) following() []Contact {
	contacts := []Contact{}

	s.getAllPrefixed("following/", &contacts)

	return contacts
}

func (s *Store) createFollower(contact *Contact) {
	s.saveStruct("follower/"+contact.Alias, contact)
}

func (s *Store) followers() []Contact {
	contacts := []Contact{}

	s.getAllPrefixed("follower/", &contacts)

	return contacts
}

//Contacts ...
func (s *Store) saveContact(contact *Contact) {
	mnemonics, _ := contact.DomainEntry.Address.Mnemonics()
	s.saveStruct("contacts/"+mnemonics, contact)
}

func (s *Store) contacts() []Contact {
	contacts := []Contact{}
	s.getAllPrefixed("contacts/", &contacts)
	return contacts
}

func (s *Store) getContact(mnemonics string) *Contact {
	var contact Contact
	err := s.getStruct("contacts/"+mnemonics, &contact)
	if err != nil {
		return nil
	}
	return &contact
}

func (s *Store) deleteContact(mnemonics string) {
	s.db.Update(func(txn *badger.Txn) error {
		txn.Delete([]byte("contacts/" + mnemonics))
		return nil
	})
}

//Messages ...
func (s *Store) saveMessage(msg *msgswap.StoredMessage) {
	s.saveStruct("messages/"+msg.ID, msg)
}

func (s *Store) messages() []msgswap.StoredMessage {
	storedMessages := []msgswap.StoredMessage{}
	s.getAllPrefixed("messages/", &storedMessages)
	return storedMessages
}

func (s *Store) getMessage(msgID string) *msgswap.StoredMessage {
	var msg msgswap.StoredMessage
	err := s.getStruct("messages/"+msgID, &msg)
	if err != nil {
		return nil
	}
	return &msg
}

func (s *Store) deleteMessage(msgID string) {
	s.db.Update(func(txn *badger.Txn) error {
		txn.Delete([]byte("messages/" + msgID))
		return nil
	})
}
