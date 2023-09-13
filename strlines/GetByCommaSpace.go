package strlines

import (
	"strings"

	"gitlab.com/auk-go/core/constants"
)

// GetByCommaSpace
//
// split by `, `
func GetByCommaSpace(content string) []string {
	allLines := strings.Split(
		content, constants.CommaSpace)

	return allLines
}
