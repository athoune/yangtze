package index

import (
	"fmt"
	"github.com/athoune/yangtze/pattern"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWatchFor(t *testing.T) {
	i, err := NewSimple()
	assert.Nil(t, err)
	p, err := pattern.Parse("beuha ... aussi")
	assert.Nil(t, err)
	i.AddPattern(p)
	fmt.Println(i)
}
