package concat

import (
	"gitlab.com/evatix-go/core/constants"
)

// Empty separator, empty string will be ignored
func Strings(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.EmptyStringPtr)
}

// Empty string will be ignored
func StringsUsingPipe(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.PipePtr)
}

// Empty string will be ignored
func StringsUsingComma(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.CommaPtr)
}

// Empty string will be ignored
func StringsUsingSpace(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.SpacePtr)
}

// Empty string will be ignored
func StringsUsingHyphen(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.HyphenPtr)
}

func StringsWithSeparator(
	currentStr,
	separator string,
	isSkipEmptyOrNil bool,
	contents ...string,
) string {
	return StringsArrayWithSeparator(
		&currentStr,
		&separator,
		isSkipEmptyOrNil,
		&contents)
}
