package pattern

import (
	"io"

	"github.com/athoune/yangtze/store"
	"github.com/athoune/yangtze/token"
	"github.com/bits-and-blooms/bitset"
)

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

func (p *Parser) Parse(src []byte) (*Pattern, error) {
	tokens := p.tokenizer.Tokenize(src)
	s := Pattern{
		Tokens:        make([]*Token, 0),
		bitset:        bitset.New(0),
		HasStartsWith: false,
	}
	var previous *Token
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
