package splits

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/defaultcapacity"
)

func LastByRunesMap(
	s *string,
	separatorRunesMap *map[rune]bool,
	limits int,
) *[]string {
	if s == nil || *s == "" {
		return defaultResult()
	}

	if limits == 0 {
		return core.EmptyStringsPtr()
	}

	if separatorRunesMap == nil || len(*separatorRunesMap) == 0 {
		return defaultResultWithStr(s)
	}

	runes := []rune(*s)
	runesLength := len(runes)
	defaultCapacity := defaultcapacity.Get(runesLength, limits)
	list := make(
		[]string,
		0,
		defaultCapacity)
	hasLimits := limits > -1
	runesIndex := runesLength - 1

	for ; runesIndex >= constants.Zero; runesIndex-- {
		curRune := runes[runesIndex]

		if (*separatorRunesMap)[curRune] == true {
			list = append(
				list,
				(*s)[runesIndex+1:runesLength])
			runesLength = runesIndex
			limits--
		}

		if hasLimits && limits <= 0 {
			break
		}
	}

	if runesLength > -1 {
		list = append(
			list,
			string(runes[0:runesLength]))
	}

	return &list
}

func LastByRune(
	s *string,
	runeToSplit rune,
	limits int,
) *[]string {
	if s == nil || *s == "" {
		return defaultResult()
	}

	if limits == 0 {
		return core.EmptyStringsPtr()
	}

	runes := []rune(*s)
	runesLength := len(runes)
	defaultCapacity := defaultcapacity.Get(runesLength, limits)
	list := make(
		[]string,
		0,
		defaultCapacity)
	hasLimits := limits > -1
	runesIndex := runesLength - 1

	for ; runesIndex >= constants.Zero; runesIndex-- {
		if runeToSplit == runes[runesIndex] {
			list = append(
				list,
				(*s)[runesIndex+1:runesLength])
			runesLength = runesIndex
			limits--
		}

		if hasLimits && limits <= 0 {
			break
		}
	}

	if runesLength > -1 {
		list = append(
			list,
			string(runes[0:runesLength]))
	}

	return &list
}
