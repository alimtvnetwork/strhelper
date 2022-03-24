package isstrs

import "gitlab.com/evatix-go/strhelper/isstr"

// AnyBlank Returns:
//  - true : if @lines are nil.
//  - true : if any line in lines is blank (whitespace or empty or nil)
func AnyBlank(lines ...string) bool {
	if lines == nil {
		return true
	}

	for _, line := range lines {
		if isstr.BlankPtr(&line) {
			return true
		}
	}

	return false
}
