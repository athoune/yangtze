package store

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/analyzer/simple"
	"github.com/blevesearch/bleve/registry"
	radix "github.com/hashicorp/go-immutable-radix"
	"sort"
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

type Set []uint32

// Sorted set without 0
func NewSet(words []uint32) Set {
	sort.Slice(words, func(i, j int) bool { return words[i] < words[j] })
	clean := []uint32{words[0]}
	last := words[0]
	for _, v := range words[1:len(words)] {
		if v != last {
			clean = append(clean, v)
			last = v
		}
	}
	if clean[0] == 0 {
		return clean[1:len(clean)]
	}
	return clean
}

func (s Set) Contains(other Set) bool {
	if len(other) > len(s) {
		return false
	}
	for _, o := range other {
		ok := false
		for _, a := range s {
			if a == o {
				ok = true
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
