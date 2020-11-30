package testscore

func GetAssertMessage(testCaseMessenger TestCaseMessenger, counter int) string {
	return GetAssertMessageQuick(
		testCaseMessenger.Value(),
		testCaseMessenger.Actual(),
		testCaseMessenger.Expected(),
		counter)
}
