package strhash

import "gitlab.com/auk-go/strhelper/anyto"

func SimpleHash(str string) int64 {
	allBytes := []byte(str)

	return SimpleAnyHash(&allBytes)
}

func SimpleAnyHash(any interface{}) int64 {
	// TODO : Simplify the formula in future.
	allBytes, err := anyto.Bytes(any)

	if err != nil {
		panic(err)
	}

	var length = int64(len(allBytes))
	mid := length / 2
	var sum = length + mid

	var bInt, index int64
	for i, b := range allBytes {
		bInt = int64(b)
		index = int64(i)
		sum += int64(i) + bInt - mid - (index | bInt) + (bInt & mid)
	}

	return sum
}
