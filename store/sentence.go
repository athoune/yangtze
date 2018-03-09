package store

import (
	"io"
)

type Sentence struct {
	Words []Word
}

func NewSentence(words ...Word) *Sentence {
	return &Sentence{
		Words: words,
	}
}

func (s *Store) AddSentence(sentence []byte) *Sentence {
	tokens := s.Tokenizer.Tokenize(sentence)
	r := NewSentence()
	for token, err := tokens.Read(); err != io.EOF; token, err = tokens.Read() {
		r.Words = append(r.Words, s.AddWord(token))
	}
	return r
}

const bufferSize = 64

func (s *Store) Sentence(sentence []byte) *Sentence {
	tokens := s.Tokenizer.Tokenize(sentence)
	cpt := 0
	r := &Sentence{
		Words: make([]Word, bufferSize),
	}
	for tok, err := tokens.Read(); err != io.EOF; tok, err = tokens.Read() {
		if cpt < bufferSize {
			r.Words[cpt] = s.Word(tok)
		} else {
			r.Words = append(r.Words, s.Word(tok))
		}
		cpt += 1
	}
	r.Words = r.Words[:cpt]
	return r
}

func (s *Sentence) Add(word Word) {
	s.Words = append(s.Words, word)
}

func (s *Sentence) Length() int {
	return len(s.Words)
}

func (s *Sentence) Slice(start int, end int) *Sentence {
	return &Sentence{
		Words: s.Words[start:end],
	}
}

func (s *Sentence) Index(substr *Sentence) int {
	return Index(s.Words, substr.Words)
}

func Index(s []Word, substr []Word) int {
	if len(substr) == 0 || len(s) == 0 || len(substr) > len(s) {
		return -1
	}
	i2 := 0
	for i1, w1 := range s {
		if w1 == substr[i2] {
			i2 += 1
			if i2 == len(substr) {
				return i1 - len(substr) + 1
			}
		} else {
			i2 = 0
		}
	}
	return -1
}
