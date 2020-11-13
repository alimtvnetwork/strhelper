package strhelper

import "gitlab.com/evatix-go/strhelper/charhelper"

// Returns
//  - Upper case string as pointer of string
//  - invalid case: the same pointer back if nil
//
// Warning: It requires more memory to copy and then case string also the order is BigO(n)
func ToUpperStrPtr(s *string) *string {
	if s == nil {
		// return as is.
		return s
	}

	runes := []rune(*s)
	runes = *charhelper.ToUpperRunesInPlace(&runes)
	toUpperString := string(runes)

	return &toUpperString
}
