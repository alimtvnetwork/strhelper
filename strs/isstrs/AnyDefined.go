package isstrs

import (
	"gitlab.com/evatix-go/strhelper/isstr"
)

// Returns:
//  - false : if @lines are nil.
//  - true : if all @lines are defined (not whitespace or empty or nil)
func AnyDefined(lines *[]string) bool {
	if lines == nil {
		return false
	}

	for _, line := range *lines {
		if isstr.DefinedPtr(&line) {
			return true
		}
	}

	return false
}
