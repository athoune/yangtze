package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStore(t *testing.T) {
	s := New()
	a1 := s.Word([]byte("Pim"))
	a2 := s.Word([]byte("Pam"))
	a3 := s.Word([]byte("Poum"))
	aa1 := s.Word([]byte("Pim"))
	assert.Equal(t, a1, uint32(1))
	assert.Equal(t, a2, uint32(2))
	assert.Equal(t, a3, uint32(3))
	assert.Equal(t, a1, aa1)
}
