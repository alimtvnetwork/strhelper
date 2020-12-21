package whitespace

func IsEmptyOrSpacesPtr(s *string) bool {
	return s == nil || *s == "" || IsWhitespaces(s)
}
