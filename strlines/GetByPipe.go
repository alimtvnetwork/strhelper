package strlines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// GetByPipe
//
// split by `|`
func GetByPipe(content string) []string {
	allLines := strings.Split(content, constants.Pipe)

	return allLines
}
