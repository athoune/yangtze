package token

import (
	"bytes"
	"io"
	"unicode"
	"unicode/utf8"
)

func DoIKeep(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_'
}

type Buffer struct {
	bytes  []byte
	buffer *bytes.Buffer
	offset int
	prems  int
}

func NewBuffer(b []byte) *Buffer {
	return &Buffer{
		bytes:  b,
		buffer: bytes.NewBuffer(b),
	}
}

func (b *Buffer) Read() ([]byte, error) {
	if b.offset == len(b.bytes) {
		return nil, io.EOF
	}
	last_is_letter := false
	for r, size, err := b.buffer.ReadRune(); true; r, size, err = b.buffer.ReadRune() {
		if err == io.EOF {
			r := b.bytes[b.prems:b.offset]
			b.prems = b.offset
			return r, nil
		}
		if err != nil {
			return nil, err
		}
		b.offset += size
		if DoIKeep(r) {
			last_is_letter = true
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

func Split(input []byte) [][]byte {
	offset := 0
	prems := 0
	out := make([][]byte, 0)
	last_is_letter := false
	for offset < len(input) {
		currRune, size := utf8.DecodeRune(input[offset:])
		offset += size
		if DoIKeep(currRune) {
			last_is_letter = true
		} else {
			if last_is_letter {
				out = append(out, input[prems:offset-size])
			}
			last_is_letter = false
			prems = offset
		}
	}
	if prems != offset {
		out = append(out, input[prems:])
	}
	return out
}
