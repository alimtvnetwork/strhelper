package strs

// s == nil || *s == nil || len(*s) == 0
func IsEmptyStringPtrArray(s *[]*string) bool {
	return s == nil || *s == nil || len(*s) == 0
}
