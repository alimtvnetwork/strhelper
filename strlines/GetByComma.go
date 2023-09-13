package strlines

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// split by `,`
func GetByComma(content *string) *[]string {
	allLines := strings.Split(*content, constants.Comma)

	return &allLines
}
