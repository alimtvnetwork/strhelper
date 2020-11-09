package brackets

type BracketStatus struct {
	IsBracketFound bool
	IsLeft         bool
	Category       Category
	FoundBracket   Bracket
	OtherBracket   Bracket
}

func EmptyBracketStatus() BracketStatus {
	return BracketStatus{
		IsBracketFound: false,
		Category:       UnknownCategory,
		FoundBracket:   UnknownBracket,
		OtherBracket:   UnknownBracket,
	}
}
