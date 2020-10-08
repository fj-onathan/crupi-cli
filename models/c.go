package models

import "encoding/binary"

// Itob: generates a Uint random Key
func Itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}
