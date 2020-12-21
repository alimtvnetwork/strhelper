package sw

func New(str string) *StringWrapper {
	stringWrapper := StringWrapper{
		content:    &str,
		runeLength: nil,
		length:     len(str),
	}

	return &stringWrapper
}

func NewPtr(str *string) *StringWrapper {
	length := 0

	if str != nil {
		length = len(*str)
	}

	stringWrapper := StringWrapper{
		content:    str,
		runeLength: nil,
		length:     length,
	}

	return &stringWrapper
}
