package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSentence(t *testing.T) {
	s1 := Sentence{1, 2, 3, 4, 5}
	s2 := Sentence{3, 4}
	assert.Equal(t, 2, s1.Index(s2))
	s3 := Sentence{5, 6}
	assert.Equal(t, -1, s1.Index(s3))
	assert.Equal(t, -1, s2.Index(s1))
}
