package constants

var (
	emptyString    = EmptyString
	space          = Space
	hyphen         = Hyphen
	comma          = Comma
	pipe           = Pipe
	newLineUnix    = NewLineUnix
	newLine        = NewLine
	tab            = Tab
	TabPtr         = &tab
	NewLinePtr     = &newLine
	EmptyStringPtr = &emptyString
	SpacePtr       = &space
	HyphenPtr      = &hyphen
	CommaPtr       = &comma
	NewLineUnixPtr = &newLineUnix

	// "|"
	PipePtr = &pipe
)

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
)
