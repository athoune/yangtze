package index

import (
	"fmt"
	"testing"
)

func TestWatchFor(t *testing.T) {
	i, err := New()
	if err != nil {
		t.Error(err)
	}
	_, err = i.WatchFor([]byte("pim pam poum pim"))
	if err != nil {
		t.Error(err)
	}

	l, err := i.Sequence([]byte("Aunt pim and the captain"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(l)
}
