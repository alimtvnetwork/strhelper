package strconcat

import (
	"encoding/json"
	"reflect"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Concatenates the inputItems and @singleContent to a single JSON string
// (@singleContent + separator + all items in inputItems).
//
//  Internally it creates a slice of interfaces pointers and
//  then use strhelpercore.AnyItems to inject and then
//  marshall from that strhelpercore.AnyItems to get the JSON result.
//
// @isSkipEmptyOrNil:
//  - Skip nil in elements.
//
// @sep:
//  - used to strconcat each strings / elements.
//
// Copied from golang library (reference : https://bit.ly/3oPHGdy).
func JsonStringExceptEmptyOfAnys(
	isSkipEmptyOrNil bool,
	singleContent *interface{},
	inputItems *[]interface{},
) *strhelpercore.StringWithError {
	items := make([]*interface{}, 0, len(*inputItems)+2)

	if isSkipEmptyOrNil && singleContent != nil && reflect.TypeOf(singleContent).Size() > 0 {
		items = append(items, singleContent)
	} else if isSkipEmptyOrNil == false {
		items = append(items, singleContent)
	}

	for _, item := range *inputItems {
		if isSkipEmptyOrNil && item == nil {
			continue
		}

		items = append(items, &item)
	}

	rawJson := strhelpercore.NewAnyItems(&items)
	jsonBytes, err := json.Marshal(rawJson)

	if err != nil {
		return strhelpercore.NewStringWithError(nil, &err)
	}

	finalStr := string(jsonBytes)

	return strhelpercore.NewStringWithNoError(&finalStr)
}
