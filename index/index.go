package index

import (
	"encoding/binary"
	"errors"
	radix "github.com/hashicorp/go-immutable-radix"
)

func encode(in uint32) []byte {
	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutUvarint(buf, uint64(in))
	return buf
}

func decode(in []byte) (uint32, error) {
	r, e := binary.Uvarint(in)
	if e == 0 {
		return 0, errors.New("buff too small")
	}
	if e < 0 {
		return 0, errors.New("buffer overflow")
	}
	return uint32(r), nil
}

type Index struct {
	words     *radix.Tree
	sequences *radix.Tree
}

func New() (*Index, error) {
	return &Index{
		words:     radix.New(),
		sequences: radix.New(),
	}, nil

}

func (i *Index) WatchFor(sequence string) error {
	return nil
}
