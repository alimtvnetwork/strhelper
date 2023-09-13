package splits

import (
	"gitlab.com/auk-go/core/constants"

	"gitlab.com/auk-go/strhelper/strhelpercore"
)

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limits (defaultResult):
//  - Given as `-1`
//
// splitStartsAt (default):
//  - Given as `0`
// isCaseSensitive (default):
//  - true
func ManyWithDefaults(
	str string,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return ManyPtr(
		str,
		splitsBy,
		0,
		constants.InvalidNotFoundCase,
		true)
}
