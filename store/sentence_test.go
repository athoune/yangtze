package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSentence(t *testing.T) {
	s1 := Sentence{1, 2, 3, 4, 5}
	s2 := Sentence{3, 4}
	assert.Equal(t, 2, s1.Index(s2))
}
