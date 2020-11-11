package brackets

type Pair struct {
	Start    Bracket
	End      Bracket
	Category Category
}

func (pair Pair) Wrap(str *string) string {
	if isEmptyStringPtr(str) {
		return pair.Start.String() + pair.End.String()
	}

	return pair.Start.String() +
		*str +
		pair.End.String()
}
