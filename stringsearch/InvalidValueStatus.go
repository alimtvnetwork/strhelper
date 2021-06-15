package stringsearch

import (
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func InvalidValueStatus(message string) *corestr.ValueStatus {
	return &corestr.ValueStatus{
		ValueValid: &corestr.ValueValid{
			Value:   constants.EmptyString,
			IsValid: false,
			Message: message,
		},
		Index: constants.InvalidNotFoundCase,
	}
}
