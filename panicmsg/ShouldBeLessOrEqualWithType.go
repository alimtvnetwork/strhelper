package panicmsg

// Returns SimpleValMsgWithType(ShouldBeLessThanEqualMessage, variableName, numberValue)
// Type name included
func ShouldBeLessOrEqualWithType(variableName string, numberValue interface{}) string {
	return SimpleValMsgWithType(ShouldBeLessThanEqualMessage, variableName, numberValue)
}
