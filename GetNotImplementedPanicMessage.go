package strhelper

import "gitlab.com/evatix-go/strhelper/constants"

func GetNotImplementedPanicMessage(url string) string {
	return constants.NotImplemented + " : [TODO] Will be solved at (" + url + ")"
}
