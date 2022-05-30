package extensions

import (
	"crypto/rand"
	"encoding/binary"
)

func NewSeed() int64 {
	b := make([]byte, 48)
	rand.Read(b)
	randNum := binary.BigEndian.Uint64(b)
	return int64(randNum)
}
