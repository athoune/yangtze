package index

import (
	"fmt"
	"testing"
)

func TestEncodeDecode(t *testing.T) {
	a := uint32(42)
	b := encode(a)
	aa, err := decode(b)
	if err != nil {
		t.Error(err)
	}
	if aa != a {
		t.Error("Oups", aa, a)
	}
}

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
