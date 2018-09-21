package pattern

import (
	"github.com/athoune/yangtze/store"
	"github.com/willf/bitset"
)

type Pattern struct {
	Tokens        []*Token
	bitset        *bitset.BitSet
	sentence      *store.Sentence
	HasStartsWith bool
}

func (p *Pattern) Match(sentence *store.Sentence) bool {
	if !sentence.Bitset.IsSuperSet(p.bitset) {
		return false
	}
	start := 0
	mode := AllStars
	for i, tok := range p.Tokens {
		switch tok.Kind {
		case Star:
			start += 1
		case JustAToken:
			idx := store.Index(sentence.Words[start:len(sentence.Words)], tok.Sentence.Words)
			if idx == -1 {
				return false
			}
			if mode == Optional && idx > 1 {
				return false
			}
			if mode == JustAToken && idx > 0 {
				return false
			}
			start += tok.Sentence.Length() + idx
			if start == sentence.Length() && (i+1) == len(p.Tokens) {
				return true
			}
		}
		mode = tok.Kind
	}
	return false
}

func (p *Pattern) Sentence() *store.Sentence {
	if p.sentence == nil {
		p.sentence = store.NewSentence()
		for _, tok := range p.Tokens {
			for _, ss := range tok.Sentence.Words {
				p.sentence.Add(ss)
			}
		}
	}
	return p.sentence
}
