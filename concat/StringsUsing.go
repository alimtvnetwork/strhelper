package concat

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

// Empty separator, empty string will be ignored
func Strings(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, strconst.EmptyStringPtr)
}

// Empty string will be ignored
func StringsUsingPipe(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, strconst.PipePtr)
}

// Empty string will be ignored
func StringsUsingComma(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, strconst.CommaPtr)
}

// Empty string will be ignored
func StringsUsingSpace(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, strconst.SpacePtr)
}

// Empty string will be ignored
func StringsUsingHyphen(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, strconst.HyphenPtr)
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
