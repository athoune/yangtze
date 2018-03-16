package index

import (
	"fmt"
	"github.com/athoune/yangtze/store"
	"github.com/athoune/yangtze/token"
	radix "github.com/hashicorp/go-immutable-radix"
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
	benchBuffer(t, b)
}

func BenchmarkBufferII(b *testing.B) {
	t := token.NewSimpleTokenizerII()
	benchBuffer(t, b)
}

func benchBuffer(t token.Tokenizer, b *testing.B) {
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
	r := make([]*regexp.Regexp, 10)
	for i := 0; i < 10; i++ {
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

func BenchmarkOneRegexp(b *testing.B) {
	r := regexp.MustCompile("beuha .* aussi")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			r.MatchString("Rien à voir")
		} else {
			r.MatchString("beuha super aussi")
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

func BenchmarkMap(b *testing.B) {
	s := map[string]int{
		"beuha": 1,
		"aussi": 2,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			_ = s["Rien"]
			_ = s["à"]
			_ = s["voir"]
		} else {
			_ = s["beuha"]
			_ = s["super"]
			_ = s["aussi"]
		}
	}
}

func BenchmarkRadix(b *testing.B) {
	s := radix.New()
	s, _, _ = s.Insert([]byte("beuha"), 1)
	s, _, _ = s.Insert([]byte("aussi"), 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			_, _ = s.Get([]byte("Rien"))
			_, _ = s.Get([]byte("à"))
			_, _ = s.Get([]byte("voir"))
		} else {
			_, _ = s.Get([]byte("beuha"))
			_, _ = s.Get([]byte("super"))
			_, _ = s.Get([]byte("aussi"))
		}
	}
}

func BenchmarkBitset(b *testing.B) {
	ok := store.NewSentence(1, 2, 3)
	ko := store.NewSentence(4, 5, 6)
	pattern := store.NewSentence(1, 2)
	for i := 0; i < b.N; i++ {
		if i%10 <= 8 {
			pattern.Bitset.IsSuperSet(ko.Bitset)
		} else {
			pattern.Bitset.IsSuperSet(ok.Bitset)
		}
	}
}
