package quotes

import "gitlab.com/evatix-go/strhelper/constants"

func isEmptyStringPtr(str *string) bool {
	return str == nil || *str == constants.EmptyString || len(*str) == 0
}
