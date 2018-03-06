package store

import (
	"io"
)

type Sentence []Word

func (s *Store) AddSentence(sentence []byte) Sentence {
	tokens := s.Tokenizer.Tokenize(sentence)
	r := make(Sentence, 0)
	for token, err := tokens.Read(); err != io.EOF; token, err = tokens.Read() {
		r = append(r, s.AddWord(token))
	}
	return r
}

const bufferSize = 64

func (s *Store) Sentence(sentence []byte) Sentence {
	tokens := s.Tokenizer.Tokenize(sentence)
	cpt := 0
	r := make(Sentence, bufferSize)
	for tok, err := tokens.Read(); err != io.EOF; tok, err = tokens.Read() {
		if cpt < bufferSize {
			r[cpt] = s.Word(tok)
		} else {
			r = append(r, s.Word(tok))
		}
		cpt += 1
	}
	return r[:cpt]
}

func (s Sentence) Index(substr Sentence) int {
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
