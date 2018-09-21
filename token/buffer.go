package token

import (
	"io"
	"unicode/utf8"
)

type Buffer struct {
	bytes  []byte
	keeper Keeper
	offset int
	prems  int
}

func (b *Buffer) Reset() {
	b.prems = 0
	b.offset = 0
}

func (b *Buffer) Read() ([]byte, error) {
	if b.offset == len(b.bytes) {
		return nil, io.EOF
	}
	last_is_letter := false
	for b.offset <= len(b.bytes) {
		r, size := utf8.DecodeRune(b.bytes[b.offset:])
		b.offset += size
		if b.keeper.DoIKeep(r) {
			last_is_letter = true
			if b.offset == len(b.bytes) {
				r := b.bytes[b.prems : b.offset-size+1]
				b.prems = b.offset
				return r, nil
			}
		} else {
			if last_is_letter {
				r := b.bytes[b.prems : b.offset-size]
				b.prems = b.offset
				return r, nil
			}
			b.prems = b.offset
		}
	}
	return nil, io.EOF
}
