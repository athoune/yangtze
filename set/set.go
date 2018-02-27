package set

import (
	"github.com/athoune/yangtze/store"
	"sort"
)

type Set []uint32

// Sorted set without 0
func NewSet(words store.Sentence) Set {
	sort.Slice(words, func(i, j int) bool { return words[i] < words[j] })
	clean := []uint32{uint32(words[0])}
	last := words[0]
	for _, v := range words[1:len(words)] {
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
