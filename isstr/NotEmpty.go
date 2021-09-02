package isstr

import "gitlab.com/evatix-go/core/constants"

func NotEmpty(s string) bool {
	return s != constants.EmptyString
}
