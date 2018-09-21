package pattern

import "github.com/athoune/yangtze/store"

type Token struct {
	Value      []byte
	Kind       Kind
	Position   int
	StartsWith bool // *
	Sentence   *store.Sentence
}

func NewToken(value []byte) *Token {
	s := len(value) > 0 && value[len(value)-1] == byte('*')
	if s {
		value = value[0 : len(value)-1]
	}
	return &Token{
		Kind:       whatKind(value),
		Value:      value,
		StartsWith: s,
		Sentence:   store.NewSentence(),
	}
}
