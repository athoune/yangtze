package token

import (
	//"github.com/stretchr/testify/assert"
	"testing"
)

func TestToken(t *testing.T) {
	tokens := Split([]byte("Beuha  aussi 42 "))
	t.Log(tokens)
	for _, tok := range tokens {
		t.Log(string(tok))
	}
}
