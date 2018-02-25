package store

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/analyzer/simple"
	"github.com/blevesearch/bleve/registry"
	radix "github.com/hashicorp/go-immutable-radix"
	"sync"
)

type Store struct {
	Words    *radix.Tree
	analyzer *analysis.Analyzer
	cache    *registry.Cache
	cpt_word uint32
	mux      sync.Mutex
}

func New(analyzer *analysis.Analyzer) *Store {
	return &Store{
		Words:    radix.New(),
		analyzer: analyzer,
		cache:    registry.NewCache(),
	}
}

func NewSimple() *Store {
	cache := registry.NewCache()
	analyzer, _ := simple.AnalyzerConstructor(nil, cache)
	return &Store{
		Words:    radix.New(),
		analyzer: analyzer,
		cache:    cache,
	}
}

func (s *Store) AddWord(word []byte) uint32 {
	k, ok := s.Words.Get(word)
	if ok {
		return k.(uint32)
	}
	s.mux.Lock()
	defer s.mux.Unlock()
	s.cpt_word += 1
	s.Words, _, _ = s.Words.Insert(word, s.cpt_word)
	return s.cpt_word
}

func (s *Store) Word(word []byte) uint32 {
	k, ok := s.Words.Get(word)
	if ok {
		return k.(uint32)
	}
	return 0
}

func (s *Store) AddSentence(sentence []byte) []uint32 {
	tokens := s.analyzer.Analyze(sentence)
	r := make([]uint32, len(tokens))
	for i, token := range tokens {
		r[i] = s.AddWord(token.Term)
	}
	return r
}

func (s *Store) Sentence(sentence []byte) []uint32 {
	tokens := s.analyzer.Analyze(sentence)
	r := make([]uint32, len(tokens))
	for i, token := range tokens {
		r[i] = s.Word(token.Term)
	}
	return r
}
