package splits

import (
	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/defaultcapacity"
)

func LastByRunePtr(
	s string,
	runeToSplit rune,
	limits int,
) []string {
	if s == "" {
		return defaultResult()
	}

	if limits == 0 {
		return []string{}
	}

	runes := []rune(s)
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

	return list
}
