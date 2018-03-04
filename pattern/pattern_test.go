package pattern

import (
	"github.com/athoune/yangtze/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {

	s := store.NewSimple()
	parser := NewParser(s)
	p, err := parser.Parse("sudo pam_unix ... session opened for user*")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(p.Tokens))
	assert.Equal(t, JustAToken, p.Tokens[0].Kind)
	assert.Equal(t, AllStars, p.Tokens[1].Kind)
	assert.Equal(t, JustAToken, p.Tokens[2].Kind)
	assert.Equal(t, 2, len(p.Tokens[0].Sentence))
	assert.Equal(t, 0, len(p.Tokens[1].Sentence))
	assert.Equal(t, 4, len(p.Tokens[2].Sentence))
}
