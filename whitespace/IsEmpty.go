package whitespace

import "gitlab.com/auk-go/core/constants"

// s == constants.EmptyString || len(s) == 0
func IsEmpty(s string) bool {
	return s == constants.EmptyString || len(s) == 0
}
