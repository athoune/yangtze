package store

type Sentence []Word

func (s *Store) AddSentence(sentence []byte) Sentence {
	tokens := s.Analyzer.Analyze(sentence)
	r := make(Sentence, len(tokens))
	for i, token := range tokens {
		r[i] = s.AddWord(token.Term)
	}
	return r
}

func (s *Store) Sentence(sentence []byte) Sentence {
	tokens := s.Analyzer.Analyze(sentence)
	r := make(Sentence, len(tokens))
	for i, token := range tokens {
		r[i] = s.Word(token.Term)
	}
	return r
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
