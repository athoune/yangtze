package token

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToken(t *testing.T) {
	tok := NewSimpleTokenizer()
	tokens := tok.Split([]byte("Beuha  aussi 42. "))
	assert.Len(t, tokens, 3)
	assert.Equal(t, "Beuha", string(tokens[0]))
	assert.Equal(t, "aussi", string(tokens[1]))
	assert.Equal(t, "42", string(tokens[2]))
}

func TestBuffer(t *testing.T) {
	tok := NewSimpleTokenizer()
	testTokenizer(tok, t)
}

func testTokenizer(tok Tokenizer, t *testing.T) {
	b := tok.Tokenize([]byte("Beuha  aussi 42"))
	l, err := b.Read()
	assert.NoError(t, err)
	assert.Equal(t, "Beuha", string(l))
	l, err = b.Read()
	assert.NoError(t, err)
	assert.Equal(t, "aussi", string(l))
	l, err = b.Read()
	assert.NoError(t, err)
	assert.Equal(t, "42", string(l))
	l, err = b.Read()
	assert.Nil(t, l)
	assert.Equal(t, io.EOF, err)
}

func TestBufferII(t *testing.T) {
	tok := NewSimpleTokenizerII()
	testTokenizer(tok, t)
}

func TestPattern(t *testing.T) {
	tok := NewSimpleTokenizer()
	b := tok.Tokenize([]byte("sudo pam_unix ... session opened for user"))
	zz := []string{"sudo", "pam_unix", "session", "opened", "for", "user"}
	for i := 0; i < 6; i++ {
		z, err := b.Read()
		assert.NoError(t, err)
		assert.Equal(t, zz[i], string(z))
	}
	_, err := b.Read()
	assert.Equal(t, io.EOF, err)
}
