package strlines

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
)

// split by `, `
func GetByCommaSpace(content *string) *[]string {
	allLines := strings.Split(*content, constants.CommaSpace)

	return &allLines
}
