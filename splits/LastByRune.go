package splits

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"

	"gitlab.com/evatix-go/strhelper/internal/defaultcapacity"
)

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
	hasLimits := limits > constants.InvalidValue
	runesIndex := runesLength - 1

	for ; runesIndex >= constants.Zero; runesIndex-- {
		if runeToSplit == runes[runesIndex] {
			list = append(
				list,
				(*s)[runesIndex+1:runesLength])
			runesLength = runesIndex
			limits--
		}

		if hasLimits && limits <= constants.One {
			break
		}
	}

	if runesLength > constants.InvalidValue {
		list = append(
			list,
			string(runes[0:runesLength]))
	}

	return &list
}
