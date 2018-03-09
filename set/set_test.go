package set

import (
	"github.com/athoune/yangtze/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	set := NewSet(store.NewSentence(2, 1, 3, 0, 0, 0))
	assert.Equal(t, set, Set{1, 2, 3})
	assert.True(t, set.Contains(Set{1, 2, 3}))
	assert.True(t, set.Contains(Set{1, 2}))
	assert.True(t, set.Contains(Set{2, 3}))
	assert.True(t, set.Contains(Set{1, 3}))
	assert.False(t, set.Contains(Set{1, 4}))
	assert.False(t, set.Contains(Set{1, 2, 3, 4}))
}
