package strcompare

import "gitlab.com/evatix-go/strhelper/isstr"

type LineEqual struct {
	IsSame                   bool
	IsSameIgnoringWhitespace bool
}

func NewLineEqual(leftLine, rightLine *Line, isCaseSensitive bool) LineEqual {
	isLineNumberSame := leftLine.Number == rightLine.Number
	isSame := leftLine == rightLine ||
		isLineNumberSame &&
			isstr.EqualsPtr(leftLine.Text(), rightLine.Text(), isCaseSensitive)
	isSameIgnoringWhitespace := isSame ||
		isLineNumberSame &&
			isstr.EqualsPtr(leftLine.TrimmedLine(), rightLine.TrimmedLine(), isCaseSensitive)

	return LineEqual{
		IsSame:                   isSame,
		IsSameIgnoringWhitespace: isSameIgnoringWhitespace,
	}
}
