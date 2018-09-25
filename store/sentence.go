package store

import (
	"io"

	"github.com/willf/bitset"
)

// Sentence is collection of Words, with a cache
type Sentence struct {
	Words  []Word
	Bitset *bitset.BitSet
}

func NewSentence(words ...Word) *Sentence {
	s := &Sentence{
		Words:  words,
		Bitset: bitset.New(uint(len(words))),
	}
	for _, word := range words {
		s.Bitset.Set(uint(word))
	}
	return s
}

func (s *Store) AddSentence(sentence []byte) *Sentence {
	tokens := s.Tokenizer.Tokenize(sentence)
	r := NewSentence()
	for token, err := tokens.Read(); err != io.EOF; token, err = tokens.Read() {
		w := s.AddWord(token)
		r.Words = append(r.Words, w)
		if w != 0 {
			r.Bitset.Set(uint(w))
		}
	}
	return r
}

const bufferSize = 64

func (s *Store) Sentence(sentence []byte) *Sentence {
	tokens := s.Tokenizer.Tokenize(sentence)
	cpt := 0
	r := &Sentence{
		Words:  make([]Word, bufferSize),
		Bitset: bitset.New(bufferSize),
	}
	for tok, err := tokens.Read(); err != io.EOF; tok, err = tokens.Read() {
		w := s.Word(tok)
		if cpt < bufferSize {
			r.Words[cpt] = w
		} else {
			r.Words = append(r.Words, w)
		}
		if w != 0 {
			r.Bitset.Set(uint(w))
		}
		cpt += 1
	}
	r.Words = r.Words[:cpt]
	return r
}

func (s *Sentence) Add(word Word) {
	s.Words = append(s.Words, word)
	if word != 0 {
		s.Bitset.Set(uint(word))
	}
}

func (s *Sentence) Length() int {
	return len(s.Words)
}

func (s *Sentence) Equal(other *Sentence) bool {
	if len(s.Words) != len(other.Words) {
		return false
	}
	for i, w := range s.Words {
		if w != other.Words[i] {
			return false
		}
	}
	return true
}

func (s *Sentence) Slice(start int, end int) *Sentence {
	return NewSentence(s.Words[start:end]...)
}

func (s *Sentence) Index(substr *Sentence) int {
	return Index(s.Words, substr.Words)
}

// Index return index of substr in s
func Index(s []Word, substr []Word) int {
	if len(substr) == 0 || len(s) == 0 || len(substr) > len(s) {
		return -1
	}
	i2 := 0
	i1 := 0
	for i1 < len(s) {
		if s[i1] == substr[i2] {
			i2++
			i1++
			if i2 == len(substr) {
				return i1 - len(substr)
			}
		} else {
			i2 = 0
			if s[i1] == substr[i2] {
				i2++
				if i2 == len(substr) {
					return i1 - len(substr)
				}
			}
			i1++
		}
	}
	return -1
}
