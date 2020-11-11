package brackets

var pairMaps = map[Category]Pair{
	Parenthesis: {
		Start:    ParenthesisStart,
		End:      ParenthesisEnd,
		Category: Parenthesis,
	},
	Curly: {
		Start:    CurlyStart,
		End:      CurlyEnd,
		Category: Curly,
	},
	Square: {
		Start:    SquareStart,
		End:      SquareEnd,
		Category: Square,
	},
}
