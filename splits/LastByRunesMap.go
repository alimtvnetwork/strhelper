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
	hasLimits := limits > constants.InvalidValue
	runesIndex := runesLength - 1

	for ; runesIndex >= constants.Zero; runesIndex-- {
		curRune := runes[runesIndex]

		if (*separatorRunesMap)[curRune] == true {
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
		selected := string(runes[0:runesLength])

		list = append(
			list,
			selected)
	}

	return &list
}
