package pattern

import (
	"github.com/athoune/yangtze/store"
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/tokenizer/whitespace"
	"strings"
)

type Kind int

const (
	JustAToken Kind = iota
	Star            // .
	AllStars        // ...
)

type Pattern struct {
	Tokens        []*Token
	HasStartsWith bool
}

type Parser struct {
	tokenizer analysis.Tokenizer
	store     *store.Store
}

func NewParser(s *store.Store) *Parser {
	tokenizer, _ := whitespace.TokenizerConstructor(nil, nil)
	return &Parser{
		tokenizer: tokenizer,
		store:     s,
	}
}

type Token struct {
	Value      string
	Kind       Kind
	Position   int
	StartsWith bool // *
	Word       store.Word
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

func (p *Parser) Parse(src string) (*Pattern, error) {
	tokens := p.tokenizer.Tokenize([]byte(src))
	s := Pattern{
		Tokens:        make([]*Token, len(tokens)),
		HasStartsWith: false,
	}
	for i, tok := range tokens {
		t := NewToken(string(tok.Term), tok.Start)
		if t.Kind == JustAToken && p.store != nil {
			t.Word = p.store.AddWord([]byte(t.Value))
		}
		s.Tokens[i] = t
		s.HasStartsWith = s.HasStartsWith || t.StartsWith
	}

	return &s, nil
}

func (p *Pattern) Sentence() store.Sentence {
	sentences := make(store.Sentence, len(p.Tokens))
	for i, t := range p.Tokens {
		if t.Kind == JustAToken {
			sentences[i] = t.Word
		} else {
			sentences[i] = 0
		}
	}
	return sentences
}

func (p *Pattern) Match(sentence store.Sentence) bool {
	return false
}
