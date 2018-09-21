package index

import (
	"sync"

	"github.com/athoune/yangtze/pattern"
	"github.com/athoune/yangtze/store"
	"github.com/athoune/yangtze/token"
)

// Index is a bag of words and an inversed index Word => Pattern
type Index struct {
	Store   *store.Store
	inverse map[store.Word][]*pattern.Pattern
	mux     sync.Mutex
}

// New Index
func New(tokenizer token.Tokenizer) (*Index, error) {
	return &Index{
		Store:   store.New(tokenizer),
		inverse: make(map[store.Word][]*pattern.Pattern),
	}, nil
}

// NewSimpleIndex return an Index, build this a SimpleStore
func NewSimpleIndex() (*Index, error) {
	return &Index{
		Store:   store.NewSimpleStore(),
		inverse: make(map[store.Word][]*pattern.Pattern),
	}, nil
}

// Parser of this Index
func (i *Index) Parser() *pattern.Parser {
	return pattern.NewParser(i.Store)
}

// AddPattern add a pattern object
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

// AddPatternBytes add a pattern line, return its Pattern object and a potential error
func (i *Index) AddPatternBytes(b []byte) (*pattern.Pattern, error) {
	p, err := i.Parser().Parse(b)
	if err != nil {
		return nil, err
	}
	i.AddPattern(p)
	return p, nil
}

// ReadLine read a line and returns patterns that match the line and if any pattern match
func (i *Index) ReadLine(line []byte) ([]*pattern.Pattern, bool) {
	sentence := i.Store.Sentence(line)
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
