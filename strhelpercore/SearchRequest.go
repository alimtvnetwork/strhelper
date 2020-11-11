package strhelpercore

// limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all
type SearchRequest struct {
	Search          string
	StartsAt        int
	// Represent how much to search for or limits the search.
	Limits          int
	IsCaseSensitive bool
}