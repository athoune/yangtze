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

func NewRadixKV() *RadixKV {
	return &RadixKV{radix.New()}
}

type MapKV struct {
	store map[string]Word
}

func (m *MapKV) Set(k []byte, v Word) {
	m.store[string(k)] = v
}

func (m *MapKV) Get(k []byte) (Word, bool) {
	v, ok := m.store[string(k)]
	return v, ok
}

func NewMapKV() *MapKV {
	return &MapKV{
		store: make(map[string]Word),
	}
}

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

func NewSimple() *Store {
	return &Store{
		Words:     NewMapKV(),
		Tokenizer: token.NewSimpleTokenizer(),
	}
}
