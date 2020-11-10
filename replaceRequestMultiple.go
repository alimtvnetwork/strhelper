package strhelper

import "gitlab.com/evatix-go/strhelper/strhelpercore"

type replaceRequestMultiple struct {
	Text *string
	// Key - represents - what to search
	//
	// Value - represents - what to replace as ReplaceRequest
	SearchReplaceMap *map[string]strhelpercore.ReplaceIndividualRequest
}
