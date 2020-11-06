package strhelper

func startAtIndexFailed(startsAtIndex int) {
	message := "startsAtIndex cannot be negative or more than the length of content. startsAtIndex:" + string(startsAtIndex)

	panic(message)
}
