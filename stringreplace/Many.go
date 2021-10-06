package stringreplace

func Many(
	text string,
	searchReplaceMap map[string]string,
	startsAt int,
	howManyReplace int, // -1 all
	isCaseSensitive bool,
) string {
	return ManyPtr(
		text,
		searchReplaceMap,
		startsAt,
		howManyReplace,
		isCaseSensitive)
}
