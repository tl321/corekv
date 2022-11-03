package codec

import "time"

type Entry struct {
	Key       []byte
	Value     []byte
	ExpiresAt uint64
}

func NewEntry(key []byte, value []byte) *Entry {
	return &Entry{
		Key:   key,
		Value: value,
	}
}

func (e *Entry) WithTTL(dur time.Duration) *Entry {
	e.ExpiresAt = uint64(time.Now().Add(dur).Unix())
	return e
}
