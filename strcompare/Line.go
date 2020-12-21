package strcompare

import (
	"strings"

	"gitlab.com/evatix-go/strhelper/isstr"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

type Line struct {
	Number      int
	text        *string
	trimmedLine *string
}

func (line *Line) Text() *string {
	return line.text
}

func (line *Line) NonNullText() string {
	if line.text == nil {
		return ""
	}

	return *line.text
}

func (line *Line) IsNull() bool {
	return line.text == nil
}

func (line *Line) IsNullOrEmpty() bool {
	return line.text == nil || *line.text == ""
}

func (line *Line) IsNullOrSpaces() bool {
	return line.text == nil || *line.text == "" || whitespace.IsWhitespaces(line.text)
}

func (line *Line) IsNullOrAsciiSpaces() bool {
	return line.text == nil || *line.text == "" || whitespace.IsAsciiWhitespaces(line.text)
}

func (line *Line) TrimmedLine() *string {
	if line.trimmedLine == nil {
		trimmedLine := strings.TrimSpace(line.NonNullText())
		line.trimmedLine = &trimmedLine
	}

	return line.trimmedLine
}

func (line *Line) Equals(anotherLine *Line, isCaseSensitive bool) LineEqual {
	return NewLineEqual(
		line,
		anotherLine,
		isCaseSensitive)
}

func (line *Line) IsSamePtr(anotherLine *Line, isCaseSensitive bool) bool {
	return line == anotherLine ||
		line.Number == anotherLine.Number &&
			isstr.EqualsPtr(line.text, anotherLine.text, isCaseSensitive)
}

func (line *Line) IsSameIgnoringWhitespace(anotherLine *Line, isCaseSensitive bool) bool {
	return line == anotherLine ||
		line.Number == anotherLine.Number &&
			isstr.EqualsPtr(line.TrimmedLine(), anotherLine.TrimmedLine(), isCaseSensitive)
}

func (line Line) IsSame(anotherLine Line, isCaseSensitive bool) bool {
	return line.Number == anotherLine.Number &&
		isstr.EqualsPtr(line.text, anotherLine.text, isCaseSensitive)
}
