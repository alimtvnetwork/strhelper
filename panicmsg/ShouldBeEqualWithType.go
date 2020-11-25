package panicmsg

// Returns SimpleValMsgWithType(ShouldBeEqualToMessage, variableName, numberValue)
// Type name included
func ShouldBeEqualWithType(variableName string, numberValue interface{}) string {
	return SimpleValMsgWithType(ShouldBeEqualToMessage, variableName, numberValue)
}
