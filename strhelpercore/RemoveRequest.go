package strhelpercore

import (
	"gitlab.com/evatix-go/core/constants"
)

type RemoveRequest struct {
	Text string
	*ReplaceIndividualRequest
}

func (removeRequest *RemoveRequest) ToReplaceRequest() *ReplaceRequest {
	return &ReplaceRequest{
		Text: removeRequest.Text,
		ReplaceIndividualRequest: &ReplaceIndividualRequest{
			Search:          removeRequest.Search,
			ReplaceWith:     constants.EmptyString,
			StartsAt:        removeRequest.StartsAt,
			HowManyReplace:  removeRequest.HowManyReplace,
			IsCaseSensitive: removeRequest.IsCaseSensitive,
			replaceCount:    nil,
		},
	}
}
