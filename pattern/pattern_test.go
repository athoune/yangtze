package pattern

import (
	"testing"
)

func TestParse(t *testing.T) {
	Parse("sudo pam_unix ... session opened for user*")
}
