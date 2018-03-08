package index

import (
	"github.com/athoune/yangtze/pattern"
	"github.com/athoune/yangtze/store"
	"github.com/athoune/yangtze/token"
	"sync"
)

type Index struct {
	store   *store.Store
	inverse map[store.Word][]*pattern.Pattern
	mux     sync.Mutex
}

func New(tokenizer token.Tokenizer) (*Index, error) {
	return &Index{
		store:   store.New(tokenizer),
		inverse: make(map[store.Word][]*pattern.Pattern),
	}, nil
}

func NewSimple() (*Index, error) {
	return &Index{
		store:   store.NewSimple(),
		inverse: make(map[store.Word][]*pattern.Pattern),
	}, nil
}

func (i *Index) Parser() *pattern.Parser {
	return pattern.NewParser(i.store)
}

func (i *Index) AddPattern(p *pattern.Pattern) {
	i.mux.Lock()
	defer i.mux.Unlock()
	for _, word := range p.Sentence() {
		if word != store.Nothing {
			if _, ok := i.inverse[word]; ok {
				i.inverse[word] = append(i.inverse[word], p)
			} else {
				i.inverse[word] = []*pattern.Pattern{p}
			}
		}
	}
}

func (i *Index) ReadLine(line []byte) ([]*pattern.Pattern, bool) {
	patterns := make([]*pattern.Pattern, 0)
	sentence := i.store.Sentence(line)
	uniq := make(map[*pattern.Pattern]bool)
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
		if p.Match(sentence) {
			patterns = append(patterns, p)
		}
	}
	return patterns, len(patterns) > 0
}
