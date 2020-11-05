package brackets

var otherBracketMaps = map[Bracket]Bracket{
	ParenthesisStart: ParenthesisEnd,
	ParenthesisEnd:   ParenthesisStart,
	CurlyStart:       CurlyEnd,
	CurlyEnd:         CurlyStart,
	SquareStart:      SquareEnd,
	SquareEnd:        SquareStart,
}
