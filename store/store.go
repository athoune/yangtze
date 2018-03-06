package store

import (
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/analyzer/simple"
	"github.com/blevesearch/bleve/registry"
	radix "github.com/hashicorp/go-immutable-radix"
	"sync"
)

const Nothing = Word(0)

type Store struct {
	Words    *radix.Tree
	Analyzer *analysis.Analyzer
	cache    *registry.Cache
	cpt_word uint32
	mux      sync.Mutex
}

func New(analyzer *analysis.Analyzer) *Store {
	return &Store{
		Words:    radix.New(),
		Analyzer: analyzer,
		cache:    registry.NewCache(),
	}
}

func NewSimple() *Store {
	cache := registry.NewCache()
	analyzer, _ := simple.AnalyzerConstructor(nil, cache)
	return &Store{
		Words:    radix.New(),
		Analyzer: analyzer,
		cache:    cache,
	}
}
