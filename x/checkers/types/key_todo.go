package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TodoKeyPrefix is the prefix to retrieve all Todo
	TodoKeyPrefix = "Todo/value/"
)

// TodoKey returns the store key to retrieve a Todo from the index fields
func TodoKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
