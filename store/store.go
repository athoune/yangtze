package store

import (
	"sync"

	"github.com/athoune/yangtze/token"
)

const Nothing = Word(0)

type Store struct {
	Words     kv
	cpt_word  uint32
	Tokenizer token.Tokenizer
	mux       sync.Mutex
}

func New(tokenizer token.Tokenizer) *Store {
	return &Store{
		Words:     NewMapKV(),
		Tokenizer: tokenizer,
	}
}

func NewSimpleStore() *Store {
	return &Store{
		Words:     NewMapKV(),
		Tokenizer: token.NewSimpleTokenizer(),
	}
}
