package index

import (
	"fmt"
	"github.com/athoune/yangtze/store"
	"github.com/athoune/yangtze/token"
	"github.com/stretchr/testify/assert"
	"io"
	"regexp"
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
	for i := 0; i < 100; i++ {
		p, err := idx.Parser().Parse([]byte(fmt.Sprintf("beuha ... aussi%v", i)))
		if err != nil {
			panic(err)
		}
		idx.AddPattern(p)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			_, _ = idx.ReadLine([]byte("Rien à voir"))
		} else {
			_, _ = idx.ReadLine([]byte("beuha super aussi42"))
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
	r := make([]*regexp.Regexp, 100)
	for i := 0; i < 100; i++ {
		r[i] = regexp.MustCompile(fmt.Sprintf("beuha .* aussi%v", i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			for _, rr := range r {
				rr.MatchString("Rien à voir")
			}
		} else {
			for _, rr := range r {
				rr.MatchString("beuha super aussi42")
			}
		}
	}
}

func BenchmarkWord(b *testing.B) {
	s := store.NewSimple()
	s.AddWord([]byte("beuha"))
	s.AddWord([]byte("aussi"))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			s.Word([]byte("Rien"))
			s.Word([]byte("à"))
			s.Word([]byte("voir"))
		} else {
			s.Word([]byte("beuha"))
			s.Word([]byte("super"))
			s.Word([]byte("aussi"))
		}
	}
}
