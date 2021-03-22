package splits

import "gitlab.com/evatix-go/core/constants"

func defaultResultWithStr(s *string) *[]string {
	if s == nil {
		return &[]string{constants.EmptyString}
	}

	return &[]string{*s}
}
