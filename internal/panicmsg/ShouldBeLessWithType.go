package panicmsg

// ShouldBeLessWithType
//
// Returns SimpleValMsgWithType(ShouldBeLessThan, variableName, numberValue)
// Type name included
func ShouldBeLessWithType(variableName string, numberValue interface{}) string {
	return SimpleValMsgWithType(ShouldBeLessThanMessage, variableName, numberValue)
}
