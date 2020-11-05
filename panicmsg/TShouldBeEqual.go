package panicmsg

// Returns TSimpleValMsg(ShouldBeEqualToMessage, variableName, numberValue)
// Type name included
func TShouldBeEqual(variableName string, numberValue interface{}) string {
	return TSimpleValMsg(ShouldBeEqualToMessage, variableName, numberValue)
}
