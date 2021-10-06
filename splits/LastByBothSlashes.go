package splits

func LastByBothSlashes(
	s string,
	limits int,
) []string {
	return LastByRunesMap(
		s,
		bothSlashesMap,
		limits)
}
