package store

import (
	"github.com/athoune/yangtze/token"
	radix "github.com/hashicorp/go-immutable-radix"
	"sync"
)

const Nothing = Word(0)

type Store struct {
	Words     *radix.Tree
	cpt_word  uint32
	Tokenizer token.Tokenizer
	mux       sync.Mutex
}

func NewSimple() *Store {
	return &Store{
		Words:     radix.New(),
		Tokenizer: token.NewSimpleTokenizer(),
	}
}
