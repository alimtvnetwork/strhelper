package panicmsg

// ShouldBeEqual
//
// Returns SimpleValMsg(ShouldBeEqualToMessage, variableName, numberValue)
func ShouldBeEqual(variableName string, numberValue int) string {
	return SimpleValMsg(ShouldBeEqualToMessage, variableName, numberValue)
}
