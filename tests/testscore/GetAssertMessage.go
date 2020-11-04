package testscore

func GetAssertMessage(testCaseMessager TestCaseMessager, counter int) string {
	return GetAssertMessageQuick(
		testCaseMessager.Value(),
		testCaseMessager.Actual(),
		testCaseMessager.Expected(),
		counter)
}
