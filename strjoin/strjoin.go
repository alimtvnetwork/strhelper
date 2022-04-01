package strjoin

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/corecsv"
	"gitlab.com/evatix-go/core/simplewrap"
)

func Apply(joiner string, lines ...string) string {
	return strings.Join(lines, joiner)
}

func JoinerSuffixMust(suffixJoiner string, lines ...string) string {
	joined := strings.Join(lines, suffixJoiner)

	if strings.HasSuffix(joined, suffixJoiner) {
		return joined
	}

	return joined + suffixJoiner
}

func JoinerPrefixMust(prefixJoiner string, lines ...string) string {
	joined := strings.Join(lines, prefixJoiner)

	if strings.HasPrefix(joined, prefixJoiner) {
		return joined
	}

	return prefixJoiner + joined
}

func Space(lines ...string) string {
	return strings.Join(lines, constants.Space)
}

// Line
//
// joined using constants.DefaultLine
func Line(lines ...string) string {
	return strings.Join(lines, constants.DefaultLine)
}

// LineEof
//
//  make sure that end of the line contains new line
func LineEof(lines ...string) string {
	joined := strings.Join(
		lines,
		constants.DefaultLine)

	if strings.HasSuffix(joined, constants.DefaultLine) {
		return joined
	}

	return joined + constants.DefaultLine
}

func Comma(lines ...string) string {
	return strings.Join(lines, constants.Comma)
}

func CommaSpace(lines ...string) string {
	return strings.Join(lines, constants.CommaSpace)
}

func Colon(lines ...string) string {
	return strings.Join(lines, constants.Colon)
}

func Hyphen(lines ...string) string {
	return strings.Join(lines, constants.Hyphen)
}

func SpaceHyphenSpace(lines ...string) string {
	return strings.Join(lines, constants.SpaceHyphenSpace)
}

func Pipe(lines ...string) string {
	return strings.Join(lines, constants.Pipe)
}

func Csv(lines ...string) string {
	return corecsv.StringsToStringDefault(
		lines...)
}

func CurlyHyphenJoin(lines ...string) string {
	for i, line := range lines {
		lines[i] = simplewrap.CurlyWrap(line)
	}

	return strings.Join(
		lines, constants.Hyphen)
}

// CsvOption
//
// Formats :
//  - isIncludeQuote && isIncludeSingleQuote = '%v' will be added
//  - isIncludeQuote && !isIncludeSingleQuote = "'%v'" will be added
//  - !isIncludeQuote && !isIncludeSingleQuote = %v will be added
func CsvOption(
	isIncludeQuote,
	isIncludeSingleQuote bool, // disable this will give double quote
	lines ...string,
) string {
	return corecsv.StringsToCsvString(
		constants.Comma,
		isIncludeQuote,
		isIncludeSingleQuote,
		lines...)
}
