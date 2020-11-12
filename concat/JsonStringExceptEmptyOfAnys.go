package concat

import (
	"encoding/json"
	"reflect"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Concat any object to array first and then compile all together as JSON
func JsonStringExceptEmptyOfAnys(
	isSkipEmptyOrNil bool,
	singleContent *interface{},
	item *[]interface{},
) *strhelpercore.StringWithError {
	items := make([]*interface{}, 0, len(*item)+2)

	if isSkipEmptyOrNil && singleContent != nil && reflect.TypeOf(singleContent).Size() > 0 {
		items = append(items, singleContent)
	} else if isSkipEmptyOrNil == false {
		items = append(items, singleContent)
	}

	for _, item := range *item {
		if isSkipEmptyOrNil && item == nil {
			continue
		}

		items = append(items, &item)
	}

	rawJson := strhelpercore.RawAnyItemsRequest{Items: &items}

	jsonBytes, err := json.Marshal(rawJson)

	if err != nil {
		return strhelpercore.NewStringWithError(nil, &err)
	}

	finalStr := string(jsonBytes)

	return strhelpercore.NewStringWithNoError(&finalStr)
}
