package panicmsg

// ShouldBeLessOrEqual
//
// Returns SimpleValMsg(ShouldBeLessThanEqualMessage, variableName, numberValue)
func ShouldBeLessOrEqual(variableName string, numberValue int) string {
	return SimpleValMsg(ShouldBeLessThanEqualMessage, variableName, numberValue)
}
