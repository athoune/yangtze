package token

import (
	"github.com/martingallagher/runes"
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

type SimpleKeeperII struct{}

func (s *SimpleKeeperII) DoIKeep(r rune) bool {
	return runes.IsLetter(r) || runes.IsDigit(r) || r == '-' || r == '_'
}

type AbstractTokenizer struct {
	keeper Keeper
}

func (t *AbstractTokenizer) Tokenize(b []byte) *Buffer {
	return &Buffer{
		bytes:  b,
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

func NewSimpleTokenizerII() Tokenizer {
	return &AbstractTokenizer{
		keeper: &SimpleKeeperII{},
	}
}

func NewNotSpaceTokenizer() Tokenizer {
	return &AbstractTokenizer{
		keeper: &NotSpaceKeeper{},
	}
}

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
