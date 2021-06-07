package strhelper

import (
	"strings"

	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/coredata/corestr"
)

func ExpandSliceBySplitsPtr(slice *[]string, splitters ...string) *[]string {
	length := corestr.LengthOfStrings(slice)
	if length == 0 {
		return &[]string{}
	}

	splitExpandFunc := func(line string) *[]string {
		if len(splitters) == 0 {
			return &[]string{}
		}

		newExpandedSlice := make([]string, 0, constants.Capacity8)

		for _, splitter := range splitters {
			lines := strings.Split(line, splitter)
			newExpandedSlice = append(newExpandedSlice, lines...)
		}

		return &newExpandedSlice
	}

	expandedSlicesOfSlice := ExpandSliceByFunc(slice, splitExpandFunc)

	return expandedSlicesOfSlice
}
