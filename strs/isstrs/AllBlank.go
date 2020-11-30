package isstrs

import "gitlab.com/evatix-go/strhelper/isstr"

// Returns:
//  - true : if @lines are nil.
//  - true : if all lines are blank (whitespace or empty or nil)
func AllBlank(lines *[]string) bool {
	if lines == nil {
		return true
	}

	for _, line := range *lines {
		if isstr.DefinedPtr(&line) {
			return false
		}
	}

	return true
}
