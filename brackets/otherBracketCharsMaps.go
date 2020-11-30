package brackets

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

var otherBracketCharsMaps = map[uint8]BracketStatus{
	strconst.ParenthesisStartSymbol: {
		IsBracketFound: true,
		Category:       Parenthesis,
		FoundBracket:   ParenthesisStart,
		OtherBracket:   ParenthesisEnd,
	},
	strconst.ParenthesisEndSymbol: {
		IsBracketFound: true,
		Category:       Parenthesis,
		FoundBracket:   ParenthesisEnd,
		OtherBracket:   ParenthesisStart,
	},
	strconst.CurlyStartSymbol: {
		IsBracketFound: true,
		Category:       Curly,
		FoundBracket:   CurlyStart,
		OtherBracket:   CurlyEnd,
	},
	strconst.CurlyEndSymbol: {
		IsBracketFound: true,
		Category:       Curly,
		FoundBracket:   CurlyEnd,
		OtherBracket:   CurlyStart,
	},
	strconst.SquareStartSymbol: {
		IsBracketFound: true,
		Category:       Square,
		FoundBracket:   SquareStart,
		OtherBracket:   SquareEnd,
	},
	strconst.SquareEndSymbol: {
		IsBracketFound: true,
		Category:       Square,
		FoundBracket:   SquareEnd,
		OtherBracket:   SquareStart,
	},
}
