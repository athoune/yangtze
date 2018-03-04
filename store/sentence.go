package store

type Sentence []Word

func (s *Store) AddSentence(sentence []byte) Sentence {
	tokens := s.analyzer.Analyze(sentence)
	r := make(Sentence, len(tokens))
	for i, token := range tokens {
		r[i] = s.AddWord(token.Term)
	}
	return r
}

func (s *Store) Sentence(sentence []byte) Sentence {
	tokens := s.analyzer.Analyze(sentence)
	r := make(Sentence, len(tokens))
	for i, token := range tokens {
		r[i] = s.Word(token.Term)
	}
	return r
}
