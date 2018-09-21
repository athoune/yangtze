package pattern

import (
	"testing"

	"github.com/athoune/yangtze/store"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	s := store.NewSimpleStore()
	parser := NewParser(s)
	p, err := parser.Parse([]byte("sudo pam_unix ... session opened for user*"))
	assert.Nil(t, err)
	assert.Equal(t, 3, len(p.Tokens))
	assert.Equal(t, JustAToken, p.Tokens[0].Kind)
	assert.Equal(t, AllStars, p.Tokens[1].Kind)
	assert.Equal(t, JustAToken, p.Tokens[2].Kind)
	assert.Equal(t, 2, p.Tokens[0].Sentence.Length())
	assert.Equal(t, 0, p.Tokens[1].Sentence.Length())
	assert.Equal(t, 4, p.Tokens[2].Sentence.Length())
}

func TestMatch(t *testing.T) {
	s := store.NewSimpleStore()
	parser := NewParser(s)
	p, err := parser.Parse([]byte("a b . d"))
	assert.Nil(t, err)
	assert.True(t, p.Match(s.Sentence([]byte("a b c d"))))
	assert.False(t, p.Match(s.Sentence([]byte("a b  d"))))
	p, err = parser.Parse([]byte("a b ... d"))
	assert.Nil(t, err)
	assert.True(t, p.Match(s.Sentence([]byte("a b c d"))))
	assert.True(t, p.Match(s.Sentence([]byte("a b a b d"))))
	p, err = parser.Parse([]byte("a b ? d"))
	assert.Nil(t, err)
	assert.True(t, p.Match(s.Sentence([]byte("a b c d"))))
	assert.True(t, p.Match(s.Sentence([]byte("a b d"))))
	assert.False(t, p.Match(s.Sentence([]byte("a b a b d"))))
}

func TestMoreMatch(t *testing.T) {
	s := store.NewSimpleStore()
	parser := NewParser(s)
	p, err := parser.Parse([]byte("beuha ... aussi"))
	assert.Nil(t, err)
	assert.False(t, p.Match(s.Sentence([]byte("Aussi super beuha"))))
}
