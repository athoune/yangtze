package token

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestToken(t *testing.T) {
	tokens := Split([]byte("Beuha  aussi 42 "))
	t.Log(tokens)
	for _, tok := range tokens {
		t.Log(string(tok))
	}
}

func TestBuffer(t *testing.T) {
	b := NewBuffer([]byte("Beuha  aussi 42 "))
	l, err := b.Read()
	assert.Nil(t, err)
	assert.Equal(t, []byte("Beuha"), l)
	l, err = b.Read()
	assert.Nil(t, err)
	assert.Equal(t, []byte("aussi"), l)
	l, err = b.Read()
	assert.Nil(t, err)
	assert.Equal(t, []byte("42"), l)
	l, err = b.Read()
	assert.Nil(t, l)
	assert.Equal(t, io.EOF, err)
}
