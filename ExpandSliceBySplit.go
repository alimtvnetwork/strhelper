package strhelper

import "gitlab.com/evatix-go/core/coredata/corestr"

func ExpandSliceBySplit(slice *[]string, splitter string) *[]string {
	length := corestr.LengthOfStrings(slice)
	if length == 0 {
		return &[]string{}
	}

	return ExpandSliceBySplitsPtr(slice, splitter)
}
