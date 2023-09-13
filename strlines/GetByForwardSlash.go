package strlines

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// GetByForwardSlash
//
// split by `/`
func GetByForwardSlash(content string) []string {
	allLines := strings.Split(
		content, constants.ForwardSlash)

	return allLines
}
