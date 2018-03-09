package pattern

import (
	"github.com/athoune/yangtze/store"
	"github.com/athoune/yangtze/token"
	"github.com/willf/bitset"
	"io"
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
	bitset        *bitset.BitSet
	HasStartsWith bool
}

type Parser struct {
	tokenizer token.Tokenizer
	store     *store.Store
}

func NewParser(s *store.Store) *Parser {
	return &Parser{
		tokenizer: token.NewNotSpaceTokenizer(),
		store:     s,
	}
}

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

func whatKind(value []byte) Kind {
	switch string(value) {
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

func (p *Parser) Parse(src []byte) (*Pattern, error) {
	tokens := p.tokenizer.Tokenize(src)
	s := Pattern{
		Tokens:        make([]*Token, 0),
		bitset:        bitset.New(0),
		HasStartsWith: false,
	}
	var previous *Token = nil
	for tok, err := tokens.Read(); err != io.EOF; tok, err = tokens.Read() {
		t := NewToken(tok)
		if previous == nil || t.Kind != JustAToken || (t.Kind == JustAToken && previous.Kind != JustAToken) {
			s.Tokens = append(s.Tokens, t)
			previous = t
		}
		if t.Kind == JustAToken {
			w := p.store.AddWord(tok)
			previous.Sentence.Add(w)
			s.bitset.Set(uint(w))
		}
		s.HasStartsWith = s.HasStartsWith || t.StartsWith
	}

	return &s, nil
}

func (p *Pattern) Match(sentence *store.Sentence) bool {
	if p.bitset.Len() > sentence.Bitset.Len() {
		return false
	}
	if !sentence.Bitset.IsSuperSet(p.bitset) {
		return false
	}
	start := 0
	mode := AllStars
	for i, tok := range p.Tokens {
		switch tok.Kind {
		case Star:
			start += 1
		case JustAToken:
			idx := store.Index(sentence.Words[start:len(sentence.Words)], tok.Sentence.Words)
			if idx == -1 {
				return false
			}
			if mode == Optional && idx > 1 {
				return false
			}
			if mode == JustAToken && idx > 0 {
				return false
			}
			start += tok.Sentence.Length() + idx
			if start == sentence.Length() && (i+1) == len(p.Tokens) {
				return true
			}
		}
		mode = tok.Kind
	}
	return false
}

func (p *Pattern) Sentence() *store.Sentence {
	s := store.NewSentence()
	for _, tok := range p.Tokens {
		for _, ss := range tok.Sentence.Words {
			s.Add(ss)
		}
	}
	return s
}
