package serialize

import (
	"encoding/binary"
	"errors"
)

func Encode(in uint32) []byte {
	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutUvarint(buf, uint64(in))
	return buf
}

func Decode(in []byte) (uint32, error) {
	r, e := binary.Uvarint(in)
	if e == 0 {
		return 0, errors.New("buff too small")
	}
	if e < 0 {
		return 0, errors.New("buffer overflow")
	}
	return uint32(r), nil
}
