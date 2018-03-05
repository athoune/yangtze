package index

import (
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

func (i *Index) Parser() *pattern.Parser {
	return pattern.NewParser(i.store)
}

func (i *Index) AddPattern(p *pattern.Pattern) {
	i.mux.Lock()
	defer i.mux.Unlock()
	i.patterns = append(i.patterns, p)
	for _, word := range p.Sentence() {
		if word != store.Nothing {
			if _, ok := i.inverse[word]; ok {
				i.inverse[word] = append(i.inverse[word], len(i.patterns))
			} else {
				i.inverse[word] = []int{len(i.patterns)}
			}
		}
	}
}

func (i *Index) ReadLine(line []byte) ([]*pattern.Pattern, bool) {
	patterns := make([]*pattern.Pattern, 0)
	sentence := i.store.Sentence(line)
	uniq := make(map[int]bool)
	for _, word := range sentence {
		if word != store.Nothing {
			for _, ps := range i.inverse[word] {
				uniq[ps] = true
			}
		}
	}
	if len(uniq) == 0 {
		return patterns, false
	}
	for p, _ := range uniq {
		pp := i.patterns[p-1]
		if pp.Match(sentence) {
			patterns = append(patterns, pp)
		}
	}
	return patterns, len(patterns) > 0
}
