package splits

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/defaultcapacity"
)

func LastByRunes(
	s *string,
	limits int,
	runes ...rune,
) *[]string {
	if s == nil || *s == "" {
		return defaultResult()
	}

	if limits == 0 {
		return core.EmptyStringsPtr()
	}

	if runes == nil {
		return defaultResultWithStr(s)
	}

	length := len(runes)

	if length == 0 {
		return defaultResultWithStr(s)
	}

	runesMap := make(
		map[rune]bool, length)

	for _, r := range runes {
		runesMap[r] = true
	}

	return LastByRunesMap(
		s,
		&runesMap,
		limits)
}

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
	index := 0
	runesIndex := runesLength - 1

	for ; runesIndex >= constants.Zero; runesIndex-- {
		curRune := runes[runesIndex]

		if (*separatorRunesMap)[curRune] == true {
			list[index] = (*s)[runesIndex:runesLength]
			runesLength = index
			index++
			limits--
		}

		if hasLimits && limits <= 0 {
			break
		}
	}

	if runesLength > 0 {
		list[index] = string(runes[0 : runesIndex+1])
	}

	return &list
}
