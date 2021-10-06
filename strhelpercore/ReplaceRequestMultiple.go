package strhelpercore

type ReplaceRequestMultiple struct {
	Text string
	// Key - represents - what to search
	//
	// Value - represents - what to replace as ReplaceRequest
	SearchReplaceMap map[string]ReplaceIndividualRequest
}
