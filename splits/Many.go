package splits

import "gitlab.com/auk-go/strhelper/strhelpercore"

// Many Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limits :
//  - number of times split will be performed for all
//  - if -1 then all split will occur
//
// splitStartsAt:
//  - where split searching will start from.
func Many(
	str string,
	splitStartsAt,
	limits int,
	isCaseSensitive bool,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return ManyPtr(
		str,
		splitsBy,
		splitStartsAt,
		limits,
		isCaseSensitive)
}
