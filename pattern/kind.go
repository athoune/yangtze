package pattern

type Kind int

const (
	JustAToken Kind = iota
	Star            // .
	Optional        // ?
	AllStars        // ...
)

func whatKind(value []byte) Kind {
	switch string(value) {
	case ".":
		return Star
	case "?":
		return Optional
	case "...":
		return AllStars
	default:
		return JustAToken
	}
}
