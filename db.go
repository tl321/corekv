package corekv

import (
	"corekv/iterator"
	"corekv/utils/codec"
	"sort"
)

type COREKVAPI interface {
	Set(*codec.Entry) error
	Get([]byte) (*codec.Entry, error)
	Del([]byte) error
	NewIterator(*iterator.Options) *Iterator
	Info() *Stats
	Close() error
}

type DB struct {
	db map[string]*codec.Entry
}

var _ COREKVAPI = &DB{}

func (db *DB) Set(e *codec.Entry) error {
	db.db[string(e.Key)] = e
	return nil
}

func (db *DB) Get(key []byte) (*codec.Entry, error) {
	return db.db[string(key)], nil
}

func (db *DB) Del(key []byte) error {
	delete(db.db, string(key))
	return nil
}

func (db *DB) NewIterator(opt *iterator.Options) *Iterator {
	nPrefix := len(opt.Prefix)
	items := make([]*codec.Entry, len(db.db))
	i := 0
	for k, v := range db.db {
		if len(k) >= nPrefix && k[:nPrefix] == string(opt.Prefix) {
			items[i] = v
			i++
		}
	}
	items = items[:i]

	if opt.IsAsc {
		sort.Slice(items, func(i, j int) bool {
			return string(items[i].Key) < string(items[j].Key)
		})
	} else {
		sort.Slice(items, func(i, j int) bool {
			return string(items[i].Key) > string(items[j].Key)
		})
	}

	return &Iterator{
		Items: items,
		Ind:   0,
		Len:   len(items),
	}
}

type Stats struct {
	EntryNum int
}

func (db *DB) Info() *Stats {
	return &Stats{EntryNum: len(db.db)}
}

func (db *DB) Close() error {
	return nil
}

type Options struct{}

func NewDefaultOptions() *Options {
	return &Options{}
}

func Open(opt *Options) *DB {
	return &DB{db: make(map[string]*codec.Entry)}
}
