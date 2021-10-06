package index

// Returns all indexes where findingString is found as a map[int]bool.
//
// Whereas Key = Index, Value = true.
//
// @limits:
//  - How many indexes should we search for and then stop looking further.
//  - `-1` means find all, 0 => nil
//
// Results:
//  - Invalid result can be nil if any (content == nil || findingString == nil) results nil.
//  - If no indexes found returns nil.
func OfAllAsKeyMap(
	content,
	findingString string,
	startsAtIndex int,
	limits int,
	isCaseSensitive bool,
) map[int]bool {
	indexes := OfAll(
		content,
		findingString,
		startsAtIndex,
		limits,
		isCaseSensitive)

	if len(indexes) == 0 || limits == 0 {
		return nil
	}

	resultingMap := make(map[int]bool, len(indexes))

	for _, valueIndex := range indexes {
		resultingMap[valueIndex] = true
	}

	return resultingMap
}
