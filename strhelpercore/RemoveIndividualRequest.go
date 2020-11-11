package strhelpercore

type RemoveIndividualRequest struct {
	Search          string
	StartsAt        int
	HowManyReplace  int
	IsCaseSensitive bool
	replaceCount    *int
}
