package store

import (
	radix "github.com/hashicorp/go-immutable-radix"
	"sync"
)

type Store struct {
	words    *radix.Tree
	cpt_word uint32
	mux      sync.Mutex
}

func New() *Store {
	return &Store{
		words: radix.New(),
	}
}

func (s *Store) Word(word []byte) uint32 {
	k, ok := s.words.Get(word)
	if ok {
		return k.(uint32)
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	s.cpt_word += 1
	s.words, _, _ = s.words.Insert(word, s.cpt_word)
	return s.cpt_word
}
