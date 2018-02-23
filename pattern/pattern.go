package pattern

import (
	"fmt"
	"github.com/blevesearch/bleve/analysis/tokenizer/regexp"
	_regexp "regexp"
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

type Pattern struct {
}

func Parse(src string) (*Sentence, error) {
	tokenizer := regexp.NewRegexpTokenizer(_regexp.MustCompile("\\S+"))

	s := Sentence{
		Tokens:        make([]*Token, 1),
		HasStartsWith: false,
	}
	for _, tok := range tokenizer.Tokenize([]byte(src)) {
		token := Token{
			Value:    string(tok.Term),
			Position: tok.Start,
		}
		fmt.Println(token)
		s.Tokens = append(s.Tokens, &token)
	}

	return &s, nil
}
