package index

import (
	"fmt"
	"testing"
)

func TestWatchFor(t *testing.T) {
	i, err := NewSimple()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(i)
}
