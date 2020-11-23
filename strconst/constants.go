package strconst

var (
	emptyString    = EmptyString
	space          = Space
	hyphen         = Hyphen
	comma          = Comma
	pipe           = Pipe
	newLineUnix    = NewLineUnix
	newLine        = NewLine
	tab            = Tab
	commaSpace     = CommaSpace
	TabPtr         = &tab
	NewLinePtr     = &newLine
	EmptyStringPtr = &emptyString
	SpacePtr       = &space
	HyphenPtr      = &hyphen
	CommaPtr       = &comma
	NewLineUnixPtr = &newLineUnix
	CommaSpacePtr  = &commaSpace

	// "|"
	PipePtr = &pipe
)

//goland:noinspection ALL
const (
	UpperCaseA                        = 'A'
	UpperCaseZ                        = 'Z'
	LowerCaseA                        = 'a'
	LowerCaseZ                        = 'z'
	LowerCase                         = LowerCaseA - UpperCaseA // c - 'A' + 'a' (ref: https://bit.ly/3mFnUPW) a - A = 32 also works
	UpperCase                         = UpperCaseA - LowerCaseA // c - 'a' + 'A'
	NewLineMac                        = "\n"
	NewLineUnix                       = "\n"
	NewLineWindows                    = "\r\n"
	Tab                               = "\t"
	TabV                              = "\v"
	EmptyString                       = ""
	Space                             = " "
	Hyphen                            = "-"
	Semicolon                         = ";"
	Colon                             = ":"
	Comma                             = ","
	CommaSpace                        = ", "
	SpaceColonSpace                   = " : "
	Pipe                              = "|"
	QuestionMarkSymbol                = "?"
	NilString                         = "nil"
	SprintValueFormat                 = "%v"
	SprintNumberFormat                = "%d"
	SprintFullPropertyNameValueFormat = "%#v"
	SprintPropertyNameValueFormat     = "%+v"
	SprintTypeFormat                  = "%T"
	InvalidNotFoundCase               = -1
	Zero                              = 0
	NotImplemented                    = "Not Implemented"
	SingleQuoteSymbol                 = '\''
	DoubleQuoteSymbol                 = '"'
	ParenthesisStartSymbol            = '('
	ParenthesisEndSymbol              = ')'
	CurlyStartSymbol                  = '{'
	CurlyEndSymbol                    = '}'
	SquareStartSymbol                 = '['
	SquareEndSymbol                   = ']'
	ArbitraryCapacity5                = 5
	ArbitraryCapacity2                = 2
	ArbitraryCapacity1                = 1
	LineFeedUnix                      = '\n'
	CarriageReturn                    = '\r'
	FormFeed                          = '\f'
	One                               = 1
	SpaceByte                         = ' '
	TabByte                           = '\t'
	LineFeedUnixByte                  = '\n'
	CarriageReturnByte                = '\r'
	FormFeedByte                      = '\f'
	TabVByte                          = '\v'
	MaxUnit8                          = 255
)

var (
	// Copied from golang strings
	AsciiSpace = [256]uint8{
		TabByte:            One,
		LineFeedUnixByte:   One,
		TabVByte:           One,
		FormFeedByte:       One,
		CarriageReturnByte: One,
		SpaceByte:          One,
		0x85:               One, // reference : https://bit.ly/2JWdIoj
		0xA0:               One, // reference : https://bit.ly/2JWdIoj
	}

	// FormFeed \f is also marked as newline here.
	AsciiNewLinesChars = [256]uint8{
		LineFeedUnix:   One,
		FormFeed:       One,
		CarriageReturn: One,
	}
)
