package whitespace

import "gitlab.com/evatix-go/strhelper/strconst"

// s == constants.EmptyString || len(s) == 0
func IsEmpty(s string) bool {
	return s == strconst.EmptyString || len(s) == 0
}
