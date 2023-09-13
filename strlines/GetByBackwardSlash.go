package strlines

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// split by `/`
func GetByBackwardSlash(content *string) *[]string {
	allLines := strings.Split(*content, constants.BackSlash)

	return &allLines
}
