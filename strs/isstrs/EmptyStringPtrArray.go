package isstrs

// s == nil || *s == nil || len(*s) == 0
func EmptyStringPtrArray(s *[]*string) bool {
	return s == nil || *s == nil || len(*s) == 0
}
