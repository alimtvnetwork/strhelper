package defaultcapacity

import "gitlab.com/evatix-go/core/constants"

func Get(wholeTextLength int, limits int) int {
	if limits > constants.MinusOne {
		return limits
	}

	defaultCapacity := wholeTextLength

	if wholeTextLength > constants.ArbitraryCapacity1000 {
		defaultCapacity = constants.ArbitraryCapacity100
	}

	if wholeTextLength > constants.ArbitraryCapacity250 {
		defaultCapacity = wholeTextLength / constants.N20
	} else {
		defaultCapacity = wholeTextLength / constants.N5
	}

	return defaultCapacity
}
