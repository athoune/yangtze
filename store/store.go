package store

import (
	radix "github.com/hashicorp/go-immutable-radix"
	"sync"
)

type Store struct {
	Words    *radix.Tree
	cpt_word uint32
	mux      sync.Mutex
}

func New() *Store {
	return &Store{
		Words: radix.New(),
	}
}

func (s *Store) Word(word []byte) uint32 {
	k, ok := s.Words.Get(word)
	if ok {
		return k.(uint32)
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	s.cpt_word += 1
	s.Words, _, _ = s.Words.Insert(word, s.cpt_word)
	return s.cpt_word
}
