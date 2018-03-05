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

func BenchmarkIndex(b *testing.B) {
	idx, err := NewSimple()
	if err != nil {
		panic(err)
	}
	p, err := idx.Parser().Parse("beuha ... aussi")
	if err != nil {
		panic(err)
	}
	idx.AddPattern(p)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			_, _ = idx.ReadLine([]byte("Rien à voir"))
		} else {
			_, _ = idx.ReadLine([]byte("Beuha super aussi"))
		}
	}
}
