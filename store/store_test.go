package store

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	s := NewSimpleStore()
	a1 := s.AddWord([]byte("pim"))
	a2 := s.AddWord([]byte("pam"))
	a3 := s.AddWord([]byte("poum"))
	aa1 := s.AddWord([]byte("pim"))
	assert.Equal(t, a1, Word(1))
	assert.Equal(t, a2, Word(2))
	assert.Equal(t, a3, Word(3))
	assert.Equal(t, a1, aa1)
	r := s.Sentence([]byte("pam, pim, poum and the captain"))
	assert.True(t, r.Equal(NewSentence(2, 1, 3, 0, 0, 0)))
}
