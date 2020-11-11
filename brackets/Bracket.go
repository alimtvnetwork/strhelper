package brackets

type Bracket uint8

const (
	UnknownBracket   Bracket = iota
	ParenthesisStart Bracket = '('
	ParenthesisEnd   Bracket = ')'
	CurlyStart       Bracket = '{'
	CurlyEnd         Bracket = '}'
	SquareStart      Bracket = '['
	SquareEnd        Bracket = ']'
)

func (bracket Bracket) GetTheOtherBracket() Bracket {
	other, _ := otherBracketMaps[bracket]

	return other
}

func (bracket Bracket) IsEqual(char uint8) bool {
	return bracket.Value() == char
}

func (bracket Bracket) String() string {
	return string(bracket)
}

func (bracket Bracket) Value() uint8 {
	return uint8(bracket)
}

func (bracket Bracket) ValuePtr() *uint8 {
	val := uint8(bracket)

	return &val
}
