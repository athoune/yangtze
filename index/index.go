package index

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/tokenizers/whitespace_tokenizer"
	"github.com/blevesearch/bleve/registry"
	radix "github.com/hashicorp/go-immutable-radix"
)

func encode(in uint32) []byte {
	buf := make([]byte, binary.MaxVarintLen32)
	binary.PutUvarint(buf, uint64(in))
	return buf
}

func decode(in []byte) (uint32, error) {
	r, e := binary.Uvarint(in)
	if e == 0 {
		return 0, errors.New("buff too small")
	}
	if e < 0 {
		return 0, errors.New("buffer overflow")
	}
	return uint32(r), nil
}

type Index struct {
	words     *radix.Tree
	sequences *radix.Tree
	cpt_word  uint32
	cpt_seq   uint32
	tokenizer analysis.Tokenizer
	cache     *registry.Cache
}

func New() (*Index, error) {
	cache := registry.NewCache()
	tokenizer, err := whitespace_tokenizer.TokenizerConstructor(nil, cache)
	if err != nil {
		return nil, err
	}
	return &Index{
		words:     radix.New(),
		sequences: radix.New(),
		cache:     cache,
		tokenizer: tokenizer,
	}, nil

}

func (i *Index) WatchFor(sequence []byte) error {
	for _, token := range i.tokenizer.Tokenize(sequence) {
		k, ok := i.words.Get(token.Term)
		var cpt uint32
		if ok {
			cpt = k.(uint32)
		} else {
			i.cpt_word += 1
			cpt = i.cpt_word
			i.words, _, ok = i.words.Insert(token.Term, cpt)
		}
		fmt.Println(token.Term, cpt)
	}
	return nil
}
