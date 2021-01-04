package brackets

import (
	"gitlab.com/evatix-go/strhelper/strconst"
)

func isEmptyStringPtr(str *string) bool {
	return str == nil || *str == strconst.EmptyString || *str == ""
}
