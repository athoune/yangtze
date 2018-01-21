package serialize

import (
	"testing"
)

func TestSerialize(t *testing.T) {
	a := uint32(42)
	s := Encode(a)
	aa, err := Decode(s)
	if err != nil {
		t.Error(err)
	}
	if aa != a {
		t.Errorf("Oups, %i != %i", aa, a)
	}
}
