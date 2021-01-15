package main

import (
	"fmt"
	"time"

	"github.com/Mrcrypt/5words/domain"
	"github.com/dgraph-io/badger/v2"
	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
	recpb "github.com/libp2p/go-libp2p-record/pb"
	"github.com/whyrusleeping/base32"
	"google.golang.org/protobuf/proto"
)

//PersistentDatastore ...
type PersistentDatastore struct {
}

// NewPersistentDatastore constructs a PersistentDatastore. It is _not_ thread-safe by
// default, wrap using sync.MutexWrap if you need thread safety (the answer here
// is usually yes).
func NewPersistentDatastore() (d *PersistentDatastore) {
	return &PersistentDatastore{}
}

// Put implements Datastore.Put
func (d *PersistentDatastore) Put(key ds.Key, value []byte) (err error) {
	return store.db.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry(append([]byte("dht/"), key.Bytes()...), value).WithTTL(time.Hour * 24 * 7)
		return txn.SetEntry(e)
	})
}

// Sync implements Datastore.Sync
func (d *PersistentDatastore) Sync(prefix ds.Key) error {
	return nil
}

// Get implements Datastore.Get
func (d *PersistentDatastore) Get(key ds.Key) (value []byte, err error) {
	err = store.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(append([]byte("dht/"), key.Bytes()...))
		if err != nil {
			return ds.ErrNotFound
		}
		value, err = item.ValueCopy(value)
		return err
	})

	return
}

// Has implements Datastore.Has
func (d *PersistentDatastore) Has(key ds.Key) (exists bool, err error) {
	err = store.db.View(func(txn *badger.Txn) error {
		_, err := txn.Get(append([]byte("dht/"), key.Bytes()...))
		if err != nil {
			return err
		}
		exists = true
		return nil
	})

	return
}

// GetSize implements Datastore.GetSize
func (d *PersistentDatastore) GetSize(key ds.Key) (size int, err error) {
	val, err := d.Get(key)
	if err != nil {
		return -1, ds.ErrNotFound
	}

	return len(val), nil
}

// Delete implements Datastore.Delete
func (d *PersistentDatastore) Delete(key ds.Key) (err error) {
	err = store.db.Update(func(txn *badger.Txn) error {
		txn.Delete(append([]byte("dht/"), key.Bytes()...))
		return nil
	})

	return
}

// Query implements Datastore.Query
func (d *PersistentDatastore) Query(q dsq.Query) (dsq.Results, error) {
	re := make([]dsq.Entry, 0, 0)

	values := map[ds.Key][]byte{}

	store.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		pre := []byte("dht/")
		for it.Seek(pre); it.ValidForPrefix(pre); it.Next() {
			item := it.Item()
			err := item.Value(func(v []byte) error {
				values[ds.NewKey(string(item.Key()))] = v
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})

	for k, v := range values {
		e := dsq.Entry{Key: k.String(), Size: len(v)}
		if !q.KeysOnly {
			e.Value = v
		}
		re = append(re, e)
	}
	r := dsq.ResultsWithEntries(q, re)
	r = dsq.NaiveQueryApply(q, r)
	return r, nil
}

//Batch ...
func (d *PersistentDatastore) Batch() (ds.Batch, error) {
	return ds.NewBasicBatch(d), nil
}

//Close ...
func (d *PersistentDatastore) Close() error {
	return nil
}

//PrintValues ...
func (d *PersistentDatastore) PrintValues() {

	values := map[ds.Key][]byte{}

	store.db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		pre := []byte("dht/")
		for it.Seek(pre); it.ValidForPrefix(pre); it.Next() {
			item := it.Item()
			var val []byte
			val, _ = item.ValueCopy(val)
			values[ds.NewKey(string(item.Key()))] = val
		}
		return nil
	})

	fmt.Printf("DHT VALUES:\n")
	for k, v := range values {

		rec := new(recpb.Record)
		rec.Unmarshal(v)

		var dom domain.Entry
		err := proto.Unmarshal(rec.Value, &dom)
		if err != nil {
			panic(err)
		}

		mnemonics, _ := dom.Address.Mnemonics()
		key, _ := base32.RawStdEncoding.DecodeString(k.String()[5:])
		fmt.Printf("Key: %s, Value: %+v\n", key, mnemonics)
		fmt.Printf("\tValues: %+v, Peers: %+v\n", dom.Values, dom.Peers)
		fmt.Printf("\tRSA Key: %x\n", dom.PubRSAKey)
	}
	fmt.Printf("\n")
}
