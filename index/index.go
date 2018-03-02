package index

import (
	"fmt"
	"github.com/athoune/yangtze/pattern"
	"github.com/athoune/yangtze/store"
	"github.com/blevesearch/bleve/analysis"
	"sync"
)

type Index struct {
	store    *store.Store
	patterns []*pattern.Pattern
	inverse  map[store.Word][]int
	mux      sync.Mutex
}

func New(analyzer *analysis.Analyzer) (*Index, error) {
	return &Index{
		store:    store.New(analyzer),
		patterns: make([]*pattern.Pattern, 0),
		inverse:  make(map[store.Word][]int),
	}, nil
}

func NewSimple() (*Index, error) {
	return &Index{
		store:    store.NewSimple(),
		patterns: make([]*pattern.Pattern, 0),
		inverse:  make(map[store.Word][]int),
	}, nil
}

func (i *Index) AddPattern(p *pattern.Pattern) {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.patterns = append(i.patterns, p)
	for _, word := range p.Sentence(i.store) {
		if word != store.Nothing {
			if _, ok := i.inverse[word]; ok {
				i.inverse[word] = append(i.inverse[word], len(i.patterns))
			} else {
				i.inverse[word] = []int{len(i.patterns)}
			}
		}
	}
}

func (i *Index) ReadLine(line []byte) {
	s := i.store.Sentence(line)
	fmt.Println(s)
}
