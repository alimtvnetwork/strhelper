package strlines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

func Join(joiner string, lines ...string) string {
	return strings.Join(lines, joiner)
}

func JoinSpace(lines ...string) string {
	return strings.Join(lines, constants.Space)
}

func JoinComma(lines ...string) string {
	return strings.Join(lines, constants.Comma)
}

func JoinCommaSpace(lines ...string) string {
	return strings.Join(lines, constants.CommaSpace)
}

func JoinColon(lines ...string) string {
	return strings.Join(lines, constants.Colon)
}

func JoinHyphen(lines ...string) string {
	return strings.Join(lines, constants.Hyphen)
}
