package charhelper

func IsMatchCaseInsensitivePtr(char1 *uint8, char2 *uint8) bool {
	return ToLowerPtr(char1) == ToLowerPtr(char2)
}
