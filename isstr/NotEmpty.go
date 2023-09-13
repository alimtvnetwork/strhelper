package isstr

import "gitlab.com/auk-go/core/constants"

func NotEmpty(s string) bool {
	return s != constants.EmptyString
}
