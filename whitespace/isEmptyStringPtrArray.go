package whitespace

// s == nil || *s == nil || len(*s) == 0
func isEmptyStringPtrArray(s *[]*string) bool {
	return s == nil || *s == nil || len(*s) == 0
}
