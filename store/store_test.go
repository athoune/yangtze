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
	assert.Equal(t, a1, uint32(1))
	assert.Equal(t, a2, uint32(2))
	assert.Equal(t, a3, uint32(3))
	assert.Equal(t, a1, aa1)
	r := s.Sentence([]byte("Pam, Pim, Poum and the captain"))
	assert.Equal(t, r, []uint32{2, 1, 3, 0, 0, 0})
	set := NewSet(r)
	assert.Equal(t, set, Set{1, 2, 3})
	assert.True(t, set.Contains(Set{1, 2, 3}))
	assert.True(t, set.Contains(Set{1, 2}))
	assert.True(t, set.Contains(Set{2, 3}))
	assert.True(t, set.Contains(Set{1, 3}))
	assert.False(t, set.Contains(Set{1, 4}))
	assert.False(t, set.Contains(Set{1, 2, 3, 4}))
}
