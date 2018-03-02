package index

import (
	"fmt"
	"github.com/athoune/yangtze/pattern"
	"github.com/athoune/yangtze/store"
	"github.com/blevesearch/bleve/analysis"
)

type Index struct {
	store    *store.Store
	patterns []*pattern.Pattern
}

func New(analyzer *analysis.Analyzer) (*Index, error) {
	return &Index{
		store:    store.New(analyzer),
		patterns: make([]*pattern.Pattern, 0),
	}, nil
}

func NewSimple() (*Index, error) {
	return &Index{
		store: store.NewSimple(),
	}, nil
}

func (i *Index) AddPattern(p *pattern.Pattern) {
	i.patterns = append(i.patterns, p)
}

func (i *Index) ReadLine(line []byte) {
	s := i.store.Sentence(line)
	fmt.Println(s)
}
