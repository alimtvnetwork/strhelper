package strhelpercore

type SearchRequest struct {
	Search          string
	StartsAt        int
	HowManyReplace  int
	IsCaseSensitive bool
}