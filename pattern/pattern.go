package pattern

import (
	"github.com/athoune/yangtze/store"
	"github.com/blevesearch/bleve/analysis/tokenizer/whitespace"
	"strings"
)

type Kind int

const (
	JustAToken Kind = iota
	Include
	Exclude
	Star     // .
	AllStars // ...
)

type Pattern struct {
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
		Kind:       JustAToken,
		Value:      value,
		Position:   position,
		StartsWith: s,
	}
}

func Parse(src string) (*Pattern, error) {
	tokenizer, _ := whitespace.TokenizerConstructor(nil, nil)

	s := Pattern{
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

func (p *Pattern) Sentence(s *Store) *store.Sentence {
	sentences := make(store.Sentence, len(p.Tokens))
	for i, t := range p.Tokens {
		if t.Kind == JustAToken {
			w := s.AddWord([]byte(t.Value))
			sentences[i] = w
		} else {
			sentences[i] = 0
		}
	}
	return &sentences
}
