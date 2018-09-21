package token

import (
	"unicode"
	"unicode/utf8"

	"github.com/martingallagher/runes"
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
