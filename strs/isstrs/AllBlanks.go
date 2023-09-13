package isstrs

import "gitlab.com/auk-go/strhelper/isstr"

// AllBlanks Returns:
//  - true : if @lines are nil.
//  - true : if all lines are blank (whitespace or empty or nil)
func AllBlanks(lines ...string) bool {
	if lines == nil {
		return true
	}

	for _, line := range lines {
		if isstr.DefinedPtr(&line) {
			return false
		}
	}

	return true
}
