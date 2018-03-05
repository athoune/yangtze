package index

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWatchFor(t *testing.T) {
	i, err := NewSimple()
	assert.Nil(t, err)
	p, err := i.Parser().Parse("beuha ... aussi")
	assert.Nil(t, err)
	i.AddPattern(p)
	_, ok := i.ReadLine([]byte("Rien à voir"))
	assert.False(t, ok)
	_, ok = i.ReadLine([]byte("Aussi super beuha"))
	assert.False(t, ok)
}
