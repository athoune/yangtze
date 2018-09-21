package pattern

import (
	"testing"

	"github.com/athoune/yangtze/store"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	s := store.NewSimpleStore()
	parser := NewParser(s)
	p, err := parser.Parse([]byte("sudo pam_unix ... session opened for user*"))
	assert.NoError(t, err)
	assert.Len(t, p.Tokens, 3)
	assert.Equal(t, JustAToken, p.Tokens[0].Kind)
	assert.Equal(t, AllStars, p.Tokens[1].Kind)
	assert.Equal(t, JustAToken, p.Tokens[2].Kind)
	assert.Equal(t, 2, p.Tokens[0].Sentence.Length())
	assert.Equal(t, 0, p.Tokens[1].Sentence.Length())
	assert.Equal(t, 4, p.Tokens[2].Sentence.Length())
}

type testPattern struct {
	parser *Parser
	store  *store.Store
	t      *testing.T
}

func newTestPattern(t *testing.T) *testPattern {
	s := store.NewSimpleStore()
	return &testPattern{
		parser: NewParser(s),
		store:  s,
		t:      t,
	}
}

func (t *testPattern) test(pattern string, lines ...interface{}) {
	p, err := t.parser.Parse([]byte(pattern))
	assert.NoError(t.t, err)
	var (
		test bool
		line string
	)
	for i := 0; i < len(lines); i += 2 {
		test = lines[i].(bool)
		line = lines[i+1].(string)
		assert.Equal(t.t, test, p.Match(t.store.Sentence([]byte(line))), "'%s' -> '%s' %v", pattern, line, test)
	}
}

func TestMatch(t *testing.T) {
	tp := newTestPattern(t)
	tp.test("a b . d",
		true, "a b c d",
		false, "a b  d",
	)
	tp.test("a b ... d",
		true, "a b c d",
		true, "a b a b d",
	)
	tp.test("a b ? d",
		true, "a b c d",
		true, "a b d",
		false, "a b a b d",
	)
	tp.test("a b ...",
		false, "b c d",
		true, "a b c d",
	)
}

func TestMoreMatch(t *testing.T) {
	s := store.NewSimpleStore()
	parser := NewParser(s)
	p, err := parser.Parse([]byte("beuha ... aussi"))
	assert.NoError(t, err)
	assert.False(t, p.Match(s.Sentence([]byte("Aussi super beuha"))))
}
