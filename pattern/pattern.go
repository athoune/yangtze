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
	Optional        // ?
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
	Sentence   store.Sentence
}

func NewToken(value string, position int) *Token {
	s := strings.HasSuffix(value, "*")
	if s {
		value = value[0 : len(value)-1]
	}
	return &Token{
		Kind:       whatKind(value),
		Value:      value,
		Position:   position,
		StartsWith: s,
		Sentence:   make(store.Sentence, 0),
	}
}

func whatKind(value string) Kind {
	switch value {
	case ".":
		return Star
	case "?":
		return Optional
	case "...":
		return AllStars
	default:
		return JustAToken
	}
}

func (p *Parser) Parse(src string) (*Pattern, error) {
	tokens := p.tokenizer.Tokenize([]byte(src))
	s := Pattern{
		Tokens:        make([]*Token, 0),
		HasStartsWith: false,
	}
	var previous *Token = nil
	for _, tok := range tokens {
		t := NewToken(string(tok.Term), tok.Start)
		if previous == nil || t.Kind != JustAToken || (t.Kind == JustAToken && previous.Kind != JustAToken) {
			s.Tokens = append(s.Tokens, t)
			previous = t
		}
		if t.Kind == JustAToken {
			previous.Sentence = append(previous.Sentence, p.store.AddWord(tok.Term))
		}
		s.HasStartsWith = s.HasStartsWith || t.StartsWith
	}

	return &s, nil
}

func (p *Pattern) Match(sentence store.Sentence) bool {
	start := 0
	mode := AllStars
	for i, tok := range p.Tokens {
		switch tok.Kind {
		case Star:
			start += 1
		case JustAToken:
			idx := sentence[start:len(sentence)].Index(tok.Sentence)
			if idx == -1 {
				return false
			}
			if mode == Optional && idx > 1 {
				return false
			}
			if mode == JustAToken && idx > 0 {
				return false
			}
			start += len(tok.Sentence) + idx
			if start == len(sentence) && (i+1) == len(p.Tokens) {
				return true
			}
		}
		mode = tok.Kind
	}
	return false
}

func (p *Pattern) Sentence() store.Sentence {
	s := make(store.Sentence, 0)
	for _, tok := range p.Tokens {
		for _, ss := range tok.Sentence {
			s = append(s, ss)
		}
	}
	return s
}
