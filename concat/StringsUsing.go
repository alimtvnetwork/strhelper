package concat

import "gitlab.com/evatix-go/strhelper/constants"

// Empty separator, empty string will be ignored
func Strings(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.EmptyStringPtr)
}

// empty string will be ignored
func StringsUsingPipe(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.PipePtr)
}

// empty string will be ignored
func StringsUsingComma(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.CommaPtr)
}

func StringsUsingSpace(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.SpacePtr)
}

func StringsUsingHyphen(contents ...string) string {
	return JoinPtrExceptEmpty(&contents, constants.HyphenPtr)
}
