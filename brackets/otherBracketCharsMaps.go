package brackets

import "gitlab.com/evatix-go/strhelper/constants"

var otherBracketCharsMaps = map[uint8]BracketStatus{
	constants.ParenthesisStartSymbol: {
		IsBracketFound: true,
		Category:       Parenthesis,
		FoundBracket:   ParenthesisStart,
		OtherBracket:   ParenthesisEnd,
	},
	constants.ParenthesisEndSymbol: {
		IsBracketFound: true,
		Category:       Parenthesis,
		FoundBracket:   ParenthesisEnd,
		OtherBracket:   ParenthesisStart,
	},
	constants.CurlyStartSymbol: {
		IsBracketFound: true,
		Category:       Curly,
		FoundBracket:   CurlyStart,
		OtherBracket:   CurlyEnd,
	},
	constants.CurlyEndSymbol: {
		IsBracketFound: true,
		Category:       Curly,
		FoundBracket:   CurlyEnd,
		OtherBracket:   CurlyStart,
	},
	constants.SquareStartSymbol: {
		IsBracketFound: true,
		Category:       Square,
		FoundBracket:   SquareStart,
		OtherBracket:   SquareEnd,
	},
	constants.SquareEndSymbol: {
		IsBracketFound: true,
		Category:       Square,
		FoundBracket:   SquareEnd,
		OtherBracket:   SquareStart,
	},
}
