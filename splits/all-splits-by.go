package splits

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/coredata/corestr"
	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/core/coreutils/stringutil"
)

func BySpace(s string) []string {
	return strings.Fields(s)
}

func ByComma(s string) []string {
	return strings.Split(
		s, constants.Comma)
}

func ByColon(s string) []string {
	return strings.Split(
		s, constants.Colon)
}

func ByHyphen(s string) []string {
	return strings.Split(
		s, constants.Hyphen)
}

func ByCommaSpace(s string) []string {
	return strings.Split(
		s, constants.CommaSpace)
}

func ByEqual(s string) []string {
	return strings.Split(
		s, constants.EqualSymbol)
}

func By(content, splitter string) []string {
	return strings.Split(
		content, splitter)
}

func TrimBy(
	content, splitter string,
	number int,
) []string {
	return stringslice.SplitTrimmedNonEmpty(
		content,
		splitter,
		number)
}

func TrimNonEmptyBy(
	content, splitter string,
) []string {
	lines := strings.Split(
		content,
		splitter,
	)

	return stringslice.NonWhitespaceTrimSlice(
		lines)
}

func KeyValByIf(
	isTrim bool,
	s string, splitter string,
) (key, val string) {
	if isTrim {
		return KeyValTrimBy(s, splitter)
	}

	return KeyValBy(s, splitter)
}

func KeyValBy(s string, splitter string) (key, val string) {
	return IntoTwo(s, splitter)
}

func KeyValTrimBy(s string, splitter string) (key, val string) {
	return IntoTwoTrimSpace(s, splitter)
}

func KeyValPairByAsIf(
	isTrim bool,
	s string, splitter string,
) (kayValPair corestr.KeyValuePair) {
	if isTrim {
		return KeyValPairTrimByAs(s, splitter)
	}

	return KeyValPairByAs(s, splitter)
}

func KeyValPairByAs(
	s string, splitter string,
) (kayValPair corestr.KeyValuePair) {
	l, r := stringutil.SplitLeftRight(
		s, splitter)

	return corestr.KeyValuePair{
		Key:   l,
		Value: r,
	}
}

func KeyValPairTrimByAs(
	s string, splitter string,
) (kayValPair corestr.KeyValuePair) {
	l, r := stringutil.SplitLeftRightTrimmed(
		s, splitter)

	return corestr.KeyValuePair{
		Key:   l,
		Value: r,
	}
}

func KeyValTrimByColon(s string) (key, val string) {
	return IntoTwoTrimSpace(s, constants.Colon)
}

func KeyValTrimByHyphen(s string) (key, val string) {
	return IntoTwoTrimSpace(s, constants.Hyphen)
}

func KeyValTrimByComma(s string) (key, val string) {
	return IntoTwoTrimSpace(s, constants.Comma)
}

func KeyValTrimBySpace(s string) (key, val string) {
	return IntoTwoTrimSpace(s, constants.Space)
}

func KeyValTrimByLine(s string) (key, val string) {
	return IntoTwoTrimSpace(s, constants.DefaultLine)
}

func KeyValuePairsBy(
	isTrim bool,
	content,
	lineSplitter,
	eachLineSplitter string, // for key val
) (kayValPair *corestr.KeyValueCollection) {
	if content == "" {
		return corestr.New.KeyValues.Empty()
	}

	lines := strings.Split(content, lineSplitter)
	trimmedLines := stringslice.NonWhitespaceTrimSliceIf(
		isTrim,
		lines)

	keyValuesPairs := corestr.
		New.
		KeyValues.
		Cap(len(trimmedLines))
	for _, line := range trimmedLines {
		key, val := KeyValByIf(
			isTrim,
			line,
			eachLineSplitter)

		keyValuesPairs.Add(
			key,
			val)
	}

	return keyValuesPairs
}
