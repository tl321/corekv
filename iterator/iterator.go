package iterator

import "corekv/utils/codec"

type Options struct {
	Prefix []byte
	IsAsc  bool
}

type Iterator interface {
	Next()
	Valid() bool
	Rewind()
	Item() Item
	Close() error
}

type Item interface {
	Entry() *codec.Entry
}
