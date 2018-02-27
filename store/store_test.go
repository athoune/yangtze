package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {
	s := NewSimple()
	a1 := s.AddWord([]byte("pim"))
	a2 := s.AddWord([]byte("pam"))
	a3 := s.AddWord([]byte("poum"))
	aa1 := s.AddWord([]byte("pim"))
	assert.Equal(t, a1, Word(1))
	assert.Equal(t, a2, Word(2))
	assert.Equal(t, a3, Word(3))
	assert.Equal(t, a1, aa1)
	r := s.Sentence([]byte("Pam, Pim, Poum and the captain"))
	assert.Equal(t, r, Sentence{2, 1, 3, 0, 0, 0})
}
