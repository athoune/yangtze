package index

import (
	"sync"

	"github.com/athoune/yangtze/pattern"
	"github.com/athoune/yangtze/store"
	"github.com/athoune/yangtze/token"
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

func NewSimpleIndex() (*Index, error) {
	return &Index{
		store:   store.NewSimpleStore(),
		inverse: make(map[store.Word][]*pattern.Pattern),
	}, nil
}

func (i *Index) Parser() *pattern.Parser {
	return pattern.NewParser(i.store)
}

func (i *Index) AddPattern(p *pattern.Pattern) {
	i.mux.Lock()
	defer i.mux.Unlock()
	for _, word := range p.Sentence().Words {
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
	sentence := i.store.Sentence(line)
	var w store.Word = 0
	for _, ww := range sentence.Words {
		if ww != 0 {
			w = ww
			break
		}
	}
	if w == 0 {
		return nil, false
	}
	patterns := make([]*pattern.Pattern, 0)
	for _, p := range i.inverse[w] {
		if p.Match(sentence) {
			patterns = append(patterns, p)
		}
	}
	return patterns, len(patterns) > 0
}
