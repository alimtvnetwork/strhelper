package panicmsg

// ShouldBeGreaterEqualWithType
//
// Returns SimpleValMsgWithType(ShouldBeGreaterThanEqualMessage, variableName, numberValue)
// Type name included
func ShouldBeGreaterEqualWithType(variableName string, numberValue interface{}) string {
	return SimpleValMsgWithType(ShouldBeGreaterThanEqualMessage, variableName, numberValue)
}
