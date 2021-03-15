package isstrinternal

import "gitlab.com/evatix-go/core/constants"

func IsCurrentOsUnix() bool {
	//goland:noinspection ALL
	return constants.NewLine == constants.NewLineUnix
}
