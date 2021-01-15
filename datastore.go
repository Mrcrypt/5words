package main

import (
	ds "github.com/ipfs/go-datastore"
	dsq "github.com/ipfs/go-datastore/query"
)

//MapDatastore ...
type MapDatastore struct {
	values map[ds.Key][]byte
}

// NewMapDatastore constructs a MapDatastore. It is _not_ thread-safe by
// default, wrap using sync.MutexWrap if you need thread safety (the answer here
// is usually yes).
func NewMapDatastore() (d *MapDatastore) {
	return &MapDatastore{
		values: make(map[ds.Key][]byte),
	}
}

// Put implements Datastore.Put
func (d *MapDatastore) Put(key ds.Key, value []byte) (err error) {
	d.values[key] = value
	return nil
}

// Sync implements Datastore.Sync
func (d *MapDatastore) Sync(prefix ds.Key) error {
	return nil
}

// Get implements Datastore.Get
func (d *MapDatastore) Get(key ds.Key) (value []byte, err error) {
	val, found := d.values[key]
	if !found {
		return nil, ds.ErrNotFound
	}
	return val, nil
}

// Has implements Datastore.Has
func (d *MapDatastore) Has(key ds.Key) (exists bool, err error) {
	_, found := d.values[key]
	return found, nil
}

// GetSize implements Datastore.GetSize
func (d *MapDatastore) GetSize(key ds.Key) (size int, err error) {
	if v, found := d.values[key]; found {
		return len(v), nil
	}
	return -1, ds.ErrNotFound
}

// Delete implements Datastore.Delete
func (d *MapDatastore) Delete(key ds.Key) (err error) {
	delete(d.values, key)
	return nil
}

// Query implements Datastore.Query
func (d *MapDatastore) Query(q dsq.Query) (dsq.Results, error) {
	re := make([]dsq.Entry, 0, len(d.values))
	for k, v := range d.values {
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
func (d *MapDatastore) Batch() (ds.Batch, error) {
	return ds.NewBasicBatch(d), nil
}

//Close ...
func (d *MapDatastore) Close() error {
	return nil
}
