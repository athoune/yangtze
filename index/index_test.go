package index

import (
	"github.com/athoune/yangtze/token"
	"github.com/stretchr/testify/assert"
	"io"
	"regexp"
	"strings"
	"testing"
)

func TestWatchFor(t *testing.T) {
	i, err := NewSimple()
	assert.Nil(t, err)
	p, err := i.Parser().Parse([]byte("beuha ... aussi"))
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
	p, err := idx.Parser().Parse([]byte("beuha ... aussi"))
	if err != nil {
		panic(err)
	}
	idx.AddPattern(p)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			_, _ = idx.ReadLine([]byte("Rien à voir"))
		} else {
			_, _ = idx.ReadLine([]byte("beuha super aussi"))
		}
	}
}

func BenchmarkToken(b *testing.B) {
	t := token.NewSimpleTokenizer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			t.Split([]byte("Rien à voir"))
		} else {
			t.Split([]byte("beuha super aussi"))
		}
	}
}

func BenchmarkBuffer(b *testing.B) {
	t := token.NewSimpleTokenizer()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			b := t.Tokenize([]byte("Rien à voir"))
			for _, err := b.Read(); err != io.EOF; _, err = b.Read() {
			}
		} else {
			b := t.Tokenize([]byte("beuha super aussi"))
			for _, err := b.Read(); err != io.EOF; _, err = b.Read() {
			}
		}
	}
}

func BenchmarkSentence(b *testing.B) {
	idx, _ := NewSimple()
	idx.Parser().Parse([]byte("beuha ... aussi"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			idx.store.Sentence([]byte("Rien à voir"))
		} else {
			idx.store.Sentence([]byte("beuha super aussi"))
		}
	}
}

func BenchmarkMatch(b *testing.B) {

	idx, _ := NewSimple()
	s1 := idx.store.Sentence([]byte("Rien à voir"))
	s2 := idx.store.Sentence([]byte("beuha super aussi"))
	p, _ := idx.Parser().Parse([]byte("beuha ... aussi"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			p.Match(s1)
		} else {
			p.Match(s2)
		}
	}
}

func BenchmarkRegexp(b *testing.B) {
	r := regexp.MustCompile("beuha .* aussi")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			r.MatchString(strings.ToLower("Rien à voir"))
		} else {
			r.MatchString(strings.ToLower("beuha super aussi"))
		}
	}
}
