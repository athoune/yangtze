package index

import (
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
