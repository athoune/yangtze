package token

import (
	"unicode"

	"github.com/martingallagher/runes"
)

type Keeper interface {
	DoIKeep(r rune) bool
}

type NotSpaceKeeper struct{}

func (t *NotSpaceKeeper) DoIKeep(r rune) bool {
	return !unicode.IsSpace(r)
}

type SimpleKeeper struct{}

func (s *SimpleKeeper) DoIKeep(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' || r == '_'
}

type SimpleKeeperII struct{}

func (s *SimpleKeeperII) DoIKeep(r rune) bool {
	return runes.IsLetter(r) || runes.IsDigit(r) || r == '-' || r == '_'
}
