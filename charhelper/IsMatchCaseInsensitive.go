package charhelper

func IsMatchCaseInsensitive(char1 uint8, char2 uint8) bool {
	return ToLower(char1) == ToLower(char2)
}
