package splits

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
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

func KeyValBy(s string, splitter string) (key, val string) {
	return IntoTwo(s, splitter)
}

func KeyValTrimBy(s string, splitter string) (key, val string) {
	return IntoTwoTrimSpace(s, splitter)
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
