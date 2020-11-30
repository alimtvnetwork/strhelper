package strto

import "gitlab.com/evatix-go/strhelper/chars"

// Returns
//  - Upper case string as pointer of string
//  - invalid case: the same pointer back if nil
//
// Warning: It requires more memory to copy and then case string also the order is BigO(n)
func UpperStrPtr(s *string) *string {
	if s == nil {
		// return as is.
		return s
	}

	runes := []rune(*s)
	toUpperString := string(*chars.ToUpperRunesInPlace(&runes))

	return &toUpperString
}
