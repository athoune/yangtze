package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSentence(t *testing.T) {
	s1 := NewSentence(1, 2, 3, 4, 5)
	assert.Equal(t, 5, s1.Length())
	assert.Equal(t, uint(5), s1.Bitset.Count())
	s2 := NewSentence(3, 4)
	assert.Equal(t, uint(2), s2.Bitset.Count())
	assert.Equal(t, 2, s1.Index(s2))
	s3 := NewSentence(5, 6)
	assert.Equal(t, -1, s1.Index(s3))
	assert.Equal(t, -1, s2.Index(s1))
	assert.Equal(t, 0, s1.Index(NewSentence(1)))
	s4 := NewSentence(1, 1, 2, 3)
	s5 := NewSentence(1, 2, 3)
	assert.Equal(t, 1, s4.Index(s5))
}
