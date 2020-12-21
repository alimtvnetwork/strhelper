package whitespace

func IsEmptyOrSpaces(s string) bool {
	return s == "" || IsWhitespaces(&s)
}
