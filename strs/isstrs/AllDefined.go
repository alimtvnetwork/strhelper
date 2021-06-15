package isstrs

import "gitlab.com/evatix-go/strhelper/isstr"

// AllDefined Returns:
//  - false : if @lines are nil.
//  - true : if all defined (NOT whitespace or empty or nil)
func AllDefined(lines *[]string) bool {
	if lines == nil {
		return false
	}

	for _, line := range *lines {
		if isstr.BlankPtr(&line) {
			return false
		}
	}

	return true
}
