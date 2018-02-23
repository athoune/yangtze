package index

import (
	"bytes"
	"fmt"
	"github.com/athoune/yangtze/serialize"
	"github.com/blevesearch/bleve/analysis"
	"github.com/blevesearch/bleve/analysis/tokenizer/whitespace"
	"github.com/blevesearch/bleve/registry"
	radix "github.com/hashicorp/go-immutable-radix"
)

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
	tokenizer, err := whitespace.TokenizerConstructor(nil, cache)
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

/*
Index a sequence of tokens
Returns sequence id and potential error
*/
func (i *Index) WatchFor(sequence []byte) (uint32, error) {
	tokens := i.tokenizer.Tokenize(sequence)
	var seq bytes.Buffer
	for _, token := range tokens {
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
		seq.Write(serialize.Encode(cpt))
	}
	fmt.Println(seq.Bytes())
	k, ok := i.sequences.Get(seq.Bytes())
	var ks uint32
	if ok {
		ks = k.(uint32)
	} else {
		i.cpt_seq += 1
		ks = i.cpt_seq
	}

	return ks, nil
}

/*
Transform a sequence of tokens (a line of logs), in a collection of token id, as bytes.

*/
func (i *Index) Sequence(line []byte) ([]byte, error) {
	tokens := i.tokenizer.Tokenize(line)
	var seq bytes.Buffer
	for _, token := range tokens {
		k, ok := i.words.Get(token.Term)
		var cpt uint32
		if ok {
			cpt = k.(uint32)
		} else {
			cpt = 0
		}
		fmt.Println(token.Term, cpt)
		seq.Write(serialize.Encode(cpt))
	}
	return seq.Bytes(), nil
}
