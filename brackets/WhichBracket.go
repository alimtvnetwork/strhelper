package brackets

func WhichBracket(char uint8, isLeft bool) BracketStatus {
	otherBracketStatus, has := otherBracketCharsMaps[char]

	if !has {
		return EmptyBracketStatus()
	}

	otherBracketStatus.IsLeft = isLeft

	return otherBracketStatus
}
