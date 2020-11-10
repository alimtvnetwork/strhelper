package brackets

type Category byte

const (
	UnknownCategory Category = iota
	Parenthesis
	Curly
	Square
)

func (category Category) Pair() Pair {
	pair, _ := pairMaps[category]

	return pair
}
