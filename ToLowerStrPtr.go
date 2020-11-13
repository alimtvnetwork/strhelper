package strhelper

import "gitlab.com/evatix-go/strhelper/charhelper"

// Returns
//  - lower string as pointer of string
//  - invalid case: the same pointer back if nil
//
// Warning: It requires more memory to copy and then case string also the order is BigO(n)
func ToLowerStrPtr(s *string) *string {
	if s == nil {
		// return as is
		return s
	}

	runes := []rune(*s)
	toLowerString := string(*charhelper.ToLowerRunesInPlace(&runes))

	return &toLowerString
}
