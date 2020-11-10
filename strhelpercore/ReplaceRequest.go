package strhelpercore

type ReplaceRequest struct {
	Text *string
	*ReplaceIndividualRequest
}

type ReplaceIndividualRequest struct {
	Search          string
	ReplaceWith     string
	StartsAt        int
	HowManyReplace  int
	IsCaseSensitive bool
	replaceCount    *int
}

func (replaceIndividualRequest *ReplaceIndividualRequest) SetReplaceCount(replaceCount int) {
	isHowManyReplaceAssigned := (*replaceIndividualRequest).HowManyReplace > -1
	isNeedsRestrictionOnReplace := isHowManyReplaceAssigned &&
		replaceCount > (*replaceIndividualRequest).HowManyReplace

	if isNeedsRestrictionOnReplace {
		(*replaceIndividualRequest).replaceCount =  &replaceIndividualRequest.HowManyReplace
	} else {
		(*replaceIndividualRequest).replaceCount = &replaceCount
	}
}

func (replaceIndividualRequest *ReplaceIndividualRequest) ShouldReplace() bool {
	return (*replaceIndividualRequest).replaceCount != nil && *replaceIndividualRequest.replaceCount > 0
}

func (replaceIndividualRequest *ReplaceIndividualRequest) ReplaceCountDecrease() {
	*replaceIndividualRequest.replaceCount--
}
