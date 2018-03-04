package store

type Word uint32

func (s *Store) AddWord(word []byte) Word {
	k, ok := s.Words.Get(word)
	if ok {
		return k.(Word)
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	s.cpt_word += 1
	s.Words, _, _ = s.Words.Insert(word, Word(s.cpt_word))
	return Word(s.cpt_word)
}

func (s *Store) Word(word []byte) Word {
	k, ok := s.Words.Get(word)
	if ok {
		return k.(Word)
	}
	return 0
}
