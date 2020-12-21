package swasync

import "sync"

func New(stringInput *string) *StringWrapper {
	length := 0

	if stringInput != nil && *stringInput != "" {
		length = len(*stringInput)
	}

	return &StringWrapper{
		content:             stringInput,
		trimmedSpaceContent: nil,
		lengthInBytes:       length,
		isNullOrEmpty:       nil,
		isEmptyOrWhitespace: nil,
		uint8s:              nil,
		bytes:               nil,
		runes:               nil,
		runesLength:         nil,
		lowerRunes:          nil,
		upperRunes:          nil,
		Mutex:               sync.Mutex{},
	}
}
