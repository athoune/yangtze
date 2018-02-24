package store

import (
	radix "github.com/hashicorp/go-immutable-radix"
)

type Store struct {
	words    *radix.Tree
	cpt_word uint32
}

func New() *Store {
	return &Store{
		words: radix.New(),
	}
}

func (s *Store) Word(word []byte) uint32 {
	k, ok := s.words.Get(word)
	var cpt uint32
	if ok {
		cpt = k.(uint32)
	} else {
		s.cpt_word += 1
		cpt = s.cpt_word
		s.words, _, ok = s.words.Insert(word, cpt)
	}
	return cpt
}
