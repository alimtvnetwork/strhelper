package splits

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
)

func Last(s, separator string) *[]string {
	if s == "" {
		return core.EmptyStringsPtr()
	}
	return LastByLimitPtr(
		&s,
		&separator,
		constants.MinusOne)
}
