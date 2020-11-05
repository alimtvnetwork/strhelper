package panicmsg

// Returns TSimpleValMsg(ShouldBeLessThanEqualMessage, variableName, numberValue)
// Type name included
func TShouldBeLessOrEqual(variableName string, numberValue interface{}) string {
	return TSimpleValMsg(ShouldBeLessThanEqualMessage, variableName, numberValue)
}
