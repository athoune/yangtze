package set

import (
	"github.com/athoune/yangtze/store"
	"sort"
)

type Set []uint32

// Sorted set without 0
func NewSet(words *store.Sentence) Set {
	sort.Slice(words.Words, func(i, j int) bool { return words.Words[i] < words.Words[j] })
	clean := []uint32{uint32(words.Words[0])}
	last := words.Words[0]
	for _, v := range words.Words[1:words.Length()] {
		if v != last {
			clean = append(clean, uint32(v))
			last = v
		}
	}
	if clean[0] == 0 {
		return clean[1:len(clean)]
	}
	return clean
}

func (s Set) Contains(other Set) bool {
	if len(other) > len(s) {
		return false
	}
	start := 0
	for _, o := range other {
		ok := false
		for i, a := range s[start:len(s)] {
			if a == o {
				ok = true
				start = i
				break
			}
			if a > o {
				return false
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
