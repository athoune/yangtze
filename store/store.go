package store

import (
	"github.com/athoune/yangtze/token"
	radix "github.com/hashicorp/go-immutable-radix"
	"sync"
)

const Nothing = Word(0)

type kv interface {
	Set([]byte, Word)
	Get([]byte) (Word, bool)
}

type RadixKV struct {
	store *radix.Tree
}

func (r *RadixKV) Set(k []byte, v Word) {
	r.store, _, _ = r.store.Insert(k, v)
}

func (r *RadixKV) Get(k []byte) (Word, bool) {
	v, ok := r.store.Get(k)
	if ok {
		return v.(Word), true
	}
	return Word(0), false
}

type Store struct {
	Words     kv
	cpt_word  uint32
	Tokenizer token.Tokenizer
	mux       sync.Mutex
}

func New(tokenizer token.Tokenizer) *Store {
	return &Store{
		Words:     &RadixKV{radix.New()},
		Tokenizer: tokenizer,
	}
}

func NewSimple() *Store {
	return &Store{
		Words:     &RadixKV{radix.New()},
		Tokenizer: token.NewSimpleTokenizer(),
	}
}
