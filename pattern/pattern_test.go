package pattern

import (
	"github.com/athoune/yangtze/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	p, err := Parse("sudo pam_unix ... session opened for user*")
	assert.Nil(t, err)
	assert.Equal(t, len(p.Tokens), 7)
	assert.Equal(t, p.Tokens[0].Value, "sudo")
	assert.Equal(t, p.Tokens[2].Kind, AllStars)
	assert.True(t, p.Tokens[6].StartsWith)
	s := store.NewSimple()
	sentence := p.Sentence(s)
	assert.Equal(t, len(sentence), len(p.Tokens))
	assert.Equal(t, sentence[2], store.Word(0))
}
