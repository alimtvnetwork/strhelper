package brackets

import "gitlab.com/evatix-go/strhelper/constants"

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

type Bracket uint8

const (
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

type Pair struct {
	Start    Bracket
	End      Bracket
	Category Category
}

func (pair Pair) Wrap(str *string) string {
	if isEmptyStringPtr(str) {
		return pair.Start.String() + pair.End.String()
	}

	return pair.Start.String() +
		*str +
		pair.End.String()
}

type BracketStatus struct {
	IsBracketFound bool
	IsLeft         bool
	Category       Category
	FoundBracket   Bracket
	OtherBracket   Bracket
}

func EmptyBracketStatus() BracketStatus {
	return BracketStatus{
		IsBracketFound: false,
		Category:       nil,
		FoundBracket:   nil,
		OtherBracket:   nil,
	}
}

func WhichBracket(char uint8, isLeft bool) BracketStatus {
	otherBracketStatus, has := otherBracketCharsMaps[char]

	if !has {
		return EmptyBracketStatus()
	}

	otherBracketStatus.IsLeft = isLeft

	return otherBracketStatus
}

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

var otherBracketMaps = map[Bracket]Bracket{
	ParenthesisStart: ParenthesisEnd,
	ParenthesisEnd:   ParenthesisStart,
	CurlyStart:       CurlyEnd,
	CurlyEnd:         CurlyStart,
	SquareStart:      SquareEnd,
	SquareEnd:        SquareStart,
}

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

func isEmptyStringPtr(str *string) bool {
	return str == nil || *str == constants.EmptyString || len(*str) == 0
}

// if isSkipOnExists true then if both brackets are exist then skip the process or
// if single is there then find the next one and apply (may not apply properly if missing)
func WrapWith(str string, category Category, isSkipOnExists bool) string {
	return WrapWithPtr(&str, category, isSkipOnExists)
}

// if isSkipOnExists true then if both brackets are exist then skip the process or
// if single is there then find the next one and apply (may not apply properly if missing)
func WrapWithPtr(str *string, category Category, isSkipOnExists bool) string {
	if isEmptyStringPtr(str) || !isSkipOnExists {
		return category.Pair().Wrap(str)
	}

	if HasBothWrappedWithPtr(str, category) {
		// no need to modify
		return *str
	}

	singleBracketStatus := getSingleBracketStatus(str)

	if singleBracketStatus.IsBracketFound && singleBracketStatus.IsLeft {
		// no need to modify
		return *str + singleBracketStatus.OtherBracket.String()
	}

	if singleBracketStatus.IsBracketFound && !singleBracketStatus.IsLeft {
		// no need to modify
		return singleBracketStatus.OtherBracket.String() + *str
	}

	return category.Pair().Wrap(str)
}

func HasBothWrappedWithPtr(str *string, category Category) bool {
	if isEmptyStringPtr(str) || len(*str) <= 1 {
		return false
	}

	// has at least 2 chars
	length := len(*str)
	pair := category.Pair()
	firstChar := (*str)[0]
	lastChar := (*str)[length-1]

	return pair.Start.IsEqual(firstChar) && pair.End.IsEqual(lastChar)
}

// This method is called when both brackets are not found.
// So the assumptions are first both bracket is done.
// Now possibility here, one exist another not
func getSingleBracketStatus(str *string) BracketStatus {
	if isEmptyStringPtr(str) {
		return EmptyBracketStatus()
	}

	// has at least 2 chars
	length := len(*str)
	firstChar := (*str)[0]
	whichBracketFirstChar := WhichBracket(firstChar, true)

	if length == 1 || whichBracketFirstChar.IsBracketFound {
		return whichBracketFirstChar
	}

	lastChar := (*str)[length-1]

	return WhichBracket(lastChar, false)
}
