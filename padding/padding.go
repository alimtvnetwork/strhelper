package padding

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/isstr"
	"gitlab.com/evatix-go/strhelper/strconst"
)

func Pad(
	str, padding *string,
	width int,
	isLeft,
	isRight bool,
) string {
	if isstr.EmptyPtr(padding) {
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

func Left(str, padding *string, width int) string {
	if isstr.EmptyPtr(padding) {
		// nothing to pad for
		return *str
	}

	paddingCompiled := strings.Repeat(*padding, width)

	return paddingCompiled + *str
}

func Right(str, padding *string, width int) string {
	if isstr.EmptyPtr(padding) {
		// nothing to pad for
		return *str
	}

	paddingCompiled := strings.Repeat(*padding, width)

	return *str + paddingCompiled
}

func Space(str *string, width int, isLeft, isRight bool) string {
	return Pad(str, strconst.SpacePtr, width, isLeft, isRight)
}

func SpaceLeft(str *string, width int) string {
	return Left(str, strconst.SpacePtr, width)
}

func SpaceRight(str *string, width int) string {
	return Right(str, strconst.SpacePtr, width)
}

func Tab(str *string, width int, isLeft, isRight bool) string {
	return Pad(str, strconst.TabPtr, width, isLeft, isRight)
}

func TabLeft(str *string, width int) string {
	return Left(str, strconst.TabPtr, width)
}

func TabRight(str *string, width int) string {
	return Right(str, strconst.TabPtr, width)
}
