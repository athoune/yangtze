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
	var k Kind
	switch value {
	case ".":
		k = Star
	case "...":
		k = AllStars
	default:
		k = JustAToken
	}
	return &Token{
		Kind:       k,
		Value:      value,
		Position:   position,
		StartsWith: s,
	}
}

func Parse(src string) (*Pattern, error) {
	tokenizer, _ := whitespace.TokenizerConstructor(nil, nil)

	tokens := tokenizer.Tokenize([]byte(src))
	s := Pattern{
		Tokens:        make([]*Token, len(tokens)),
		HasStartsWith: false,
	}
	for i, tok := range tokens {
		t := NewToken(string(tok.Term), tok.Start)
		s.Tokens[i] = t
		s.HasStartsWith = s.HasStartsWith || t.StartsWith
	}

	return &s, nil
}

func (p *Pattern) Sentence(s *store.Store) store.Sentence {
	sentences := make(store.Sentence, len(p.Tokens))
	for i, t := range p.Tokens {
		if t.Kind == JustAToken {
			w := s.AddWord([]byte(t.Value))
			sentences[i] = w
		} else {
			sentences[i] = 0
		}
	}
	return sentences
}
