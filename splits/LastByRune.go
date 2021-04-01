package splits

import (
	"gitlab.com/evatix-go/core"
	"gitlab.com/evatix-go/core/constants"
	"gitlab.com/evatix-go/core/defaultcapacity"
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
	defaultCapacity := defaultcapacity.OfSplits(runesLength, limits)
	list := make(
		[]string,
		0,
		defaultCapacity)
	hasLimits := limits > constants.InvalidValue
	runesIndex := runesLength - 1

	for ; runesIndex >= constants.Zero; runesIndex-- {
		if runeToSplit == runes[runesIndex] {
			selectedRunes := runes[runesIndex+1 : runesLength]
			list = append(
				list,
				string(selectedRunes))
			runesLength = runesIndex
			limits--
		}

		// breaking before because the remaining ones will be at the end
		if hasLimits && limits <= constants.One {
			break
		}
	}

	if runesLength > constants.InvalidValue {
		selectedRunes := runes[0:runesLength]

		list = append(
			list,
			string(selectedRunes))
	}

	return &list
}
