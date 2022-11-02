package corekv

import (
	"corekv/iterator"
	"corekv/utils/codec"
)

type Iterator struct {
	Items []*codec.Entry
	Ind   int
	Len   int
}

var _ iterator.Iterator = &Iterator{}

func (i *Iterator) Next() {
	i.Ind++
}

func (i *Iterator) Valid() bool {
	return i.Ind < i.Len
}

func (i *Iterator) Rewind() {}

func (i *Iterator) Item() iterator.Item {
	return i
}

func (i *Iterator) Entry() *codec.Entry {
	return i.Items[i.Ind]
}

func (i *Iterator) Close() error {
	return nil
}
