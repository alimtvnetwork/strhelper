package panicmsg

// Returns TSimpleValMsg(ShouldBeLessThan, variableName, numberValue)
// Type name included
func TShouldBeLess(variableName string, numberValue interface{}) string {
	return TSimpleValMsg(ShouldBeLessThanMessage, variableName, numberValue)
}
