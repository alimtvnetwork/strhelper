package strs

func startAtIndexFailed(startsAtIndex int) {
	message := "startsAtIndex cannot be negative or more than the length of lines. startsAtIndex:" + string(startsAtIndex)

	panic(message)
}
