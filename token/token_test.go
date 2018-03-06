package token

import (
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestToken(t *testing.T) {
	tokens := Split([]byte("Beuha  aussi 42 "))
	assert.Equal(t, 3, len(tokens))
	assert.Equal(t, []byte("Beuha"), tokens[0])
	assert.Equal(t, []byte("aussi"), tokens[1])
	assert.Equal(t, []byte("42"), tokens[2])
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
