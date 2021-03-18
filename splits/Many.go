package splits

import "gitlab.com/evatix-go/strhelper/strhelpercore"

// Multiple split occur from the given array of splits.
//
// Basics of split("Hello World", " ") -> ["Hello", "World"] splitter will not be available in the result.
//
// limit :
//  - number of times split will performed for all
//  - if -1 then all split will occur
//
// splitStartsAt:
//  - where split searching will start from.
func Many(
	str string,
	splitStartsAt,
	limit int,
	isCaseSensitive bool,
	splitsBy ...string,
) *strhelpercore.SplitResultOverview {
	return ManyPtr(
		&str,
		&splitsBy,
		splitStartsAt,
		limit,
		isCaseSensitive)
}
