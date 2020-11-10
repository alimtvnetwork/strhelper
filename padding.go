package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/constants"
)

func Pad(
	str, padding *string,
	width int,
	isLeft,
	isRight bool,
) string {
	if IsEmptyPtr(padding) {
		// nothing to pad for
		return *str
	}

	paddingCompiled := strings.Repeat(*padding, width)

	if isLeft && isRight {
		return paddingCompiled + *str + paddingCompiled
	}

	if isLeft {
		return paddingCompiled + *str
	}

	// right
	return *str + paddingCompiled
}

func PadLeft(str, padding *string, width int) string {
	if IsEmptyPtr(padding) {
		// nothing to pad for
		return *str
	}

	paddingCompiled := strings.Repeat(*padding, width)

	return paddingCompiled + *str
}

func PadRight(str, padding *string, width int) string {
	if IsEmptyPtr(padding) {
		// nothing to pad for
		return *str
	}

	paddingCompiled := strings.Repeat(*padding, width)

	return *str + paddingCompiled
}

func PadSpace(str *string, width int, isLeft, isRight bool) string {
	return Pad(str, constants.SpacePtr, width, isLeft, isRight)
}

func PadSpaceLeft(str *string, width int) string {
	return PadLeft(str, constants.SpacePtr, width)
}

func PadSpaceRight(str *string, width int) string {
	return PadRight(str, constants.SpacePtr, width)
}

func PadTab(str *string, width int, isLeft, isRight bool) string {
	return Pad(str, constants.TabPtr, width, isLeft, isRight)
}

func PadTabLeft(str *string, width int) string {
	return PadLeft(str, constants.TabPtr, width)
}

func PadTabRight(str *string, width int) string {
	return PadRight(str, constants.TabPtr, width)
}
