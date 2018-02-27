package pattern

import (
	"github.com/blevesearch/bleve/analysis/tokenizer/whitespace"
	"strings"
)

type Kind int

const (
	Include Kind = iota
	Exclude
	Star     // .
	AllStars // ...
)

type Sentence struct {
	Tokens        []*Token
	HasStartsWith bool
}

type Token struct {
	Value      string
	Kind       Kind
	Position   int
	StartsWith bool // *
}

func NewToken(value string, position int) *Token {
	s := strings.HasSuffix(value, "*")
	if s {
		value = value[0 : len(value)-1]
	}
	return &Token{
		Value:      value,
		Position:   position,
		StartsWith: s,
	}
}

func Parse(src string) (*Sentence, error) {
	tokenizer, _ := whitespace.TokenizerConstructor(nil, nil)

	s := Sentence{
		Tokens:        make([]*Token, 1),
		HasStartsWith: false,
	}
	for _, tok := range tokenizer.Tokenize([]byte(src)) {
		t := NewToken(string(tok.Term), tok.Start)
		s.Tokens = append(s.Tokens, t)
		s.HasStartsWith = s.HasStartsWith || t.StartsWith
	}

	return &s, nil
}
