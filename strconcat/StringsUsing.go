package strconcat

import (
	"gitlab.com/auk-go/core/constants"
)

// Empty separator, empty string will be ignored
func Strings(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.EmptyString)
}

// Empty string will be ignored
func StringsUsingPipe(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.Pipe)
}

// Empty string will be ignored
func StringsUsingComma(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.Comma)
}

// Empty string will be ignored
func StringsUsingSpace(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.Space)
}

// Empty string will be ignored
func StringsUsingHyphen(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.Hyphen)
}

func StringsWithSeparator(
	currentStr,
	separator string,
	isSkipEmptyOrNil bool,
	contents ...string,
) string {
	return StringsArrayWithSeparator(
		currentStr,
		separator,
		isSkipEmptyOrNil,
		&contents)
}
