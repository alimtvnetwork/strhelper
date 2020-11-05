package brackets

type Category byte

const (
	Parenthesis Category = iota
	Curly
	Square
)

func (category Category) Pair() Pair {
	pair, _ := pairMaps[category]

	return pair
}
