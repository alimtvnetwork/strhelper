package strlines

import (
	"sort"

	"gitlab.com/auk-go/core/coredata/stringslice"
)

func SortClone(lines ...string) []string {
	cloned := stringslice.Clone(lines)
	sort.Strings(cloned)

	return cloned
}

func SortCloneIf(
	isSortClone bool,
	lines ...string,
) []string {
	if !isSortClone {
		return lines
	}

	cloned := stringslice.Clone(lines)
	sort.Strings(cloned)

	return cloned
}

func Clone(lines ...string) []string {
	return stringslice.Clone(lines)
}

func CloneIf(
	isClone bool,
	lines ...string,
) []string {
	if !isClone {
		return lines
	}

	return stringslice.Clone(lines)
}
