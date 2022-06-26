package header

import (
	"debug/elf"
	"errors"
	"sync"
)

var UnmarshalError = errors.New("an error occurred in Unmarshal")

type RequestHeader struct {
	sync.RWMutex
	CompressType elf.CompressionType
}
