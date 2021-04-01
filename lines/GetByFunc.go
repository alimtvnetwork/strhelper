package lines

import "strings"

// split by func
func GetByFunc(content *string, f func(rune) bool) *[]string {
	allLines := strings.FieldsFunc(*content, f)

	return &allLines
}
