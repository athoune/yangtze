package token

import (
	"bytes"
	"io"
	"unicode"
	"unicode/utf8"
)

type Keeper interface {
	DoIKeep(r rune) bool
}

type NotSpaceKeeper struct{}

func (t *NotSpaceKeeper) DoIKeep(r rune) bool {
	return !unicode.IsSpace(r)
}

type SimpleKeeper struct{}

func (s *SimpleKeeper) DoIKeep(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_'
}

type AbstractTokenizer struct {
	keeper Keeper
}

func (t *AbstractTokenizer) Tokenize(b []byte) *Buffer {
	return &Buffer{
		bytes:  b,
		buffer: bytes.NewBuffer(b),
		keeper: t.keeper,
	}
}

type Tokenizer interface {
	Tokenize(b []byte) *Buffer
	Split(input []byte) [][]byte
}

func NewSimpleTokenizer() Tokenizer {
	return &AbstractTokenizer{
		keeper: &SimpleKeeper{},
	}
}

func NewNotSpaceTokenizer() Tokenizer {
	return &AbstractTokenizer{
		keeper: &NotSpaceKeeper{},
	}
}

type Buffer struct {
	bytes  []byte
	buffer *bytes.Buffer
	keeper Keeper
	offset int
	prems  int
}

func (b *Buffer) Reset() {
	b.buffer = bytes.NewBuffer(b.bytes)
	b.prems = 0
	b.offset = 0
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
		if b.keeper.DoIKeep(r) {
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

func (t *AbstractTokenizer) Split(input []byte) [][]byte {
	offset := 0
	prems := 0
	out := make([][]byte, 0)
	last_is_letter := false
	for offset < len(input) {
		currRune, size := utf8.DecodeRune(input[offset:])
		offset += size
		if t.keeper.DoIKeep(currRune) {
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
