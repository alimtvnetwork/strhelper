package swp

func New(stringInput *string) *StringWrapper {
	length := 0

	if stringInput != nil && *stringInput != "" {
		length = len(*stringInput)
	}

	return &StringWrapper{
		content:             stringInput,
		trimmedSpaceContent: nil,
		lengthInBytes:       length,
		uint8s:              nil,
		bytes:               nil,
		runes:               nil,
		runesLength:         nil,
		lowerRunes:          nil,
		upperRunes:          nil,
	}
}
