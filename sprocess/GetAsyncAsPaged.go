package sprocess

import (
	"math"
	"sync"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Runs loop in async mode (in golang starts with go).
//
// It requires more memory to deal with parallel execution.
//
// Sometimes it is faster for large collection in async, memory is cheap than idle cpu.
//
// It works in multiple pages, should be faster than GetAsync
// @pages:
//  - Indicates how many item to be process in single page parallel.
func GetAsyncAsPaged(
	content *interface{},
	allRawItems *[]interface{},
	genericProcessor *strhelpercore.GenericProcessor,
	pages float64,
) *[]*interface{} {
	length := len(*allRawItems)
	processedItems := make([]*interface{}, length)
	var wg sync.WaitGroup
	pagesInInt := int(pages)
	wg.Add(pagesInInt)

	// https://bit.ly/35H47Ku
	// val skipItems = eachPageItems * (pageIndex - 1)

	eachPageItems := int(math.Ceil(float64(length) / pages))

	for pageIndex := 0; pageIndex < pagesInInt; pageIndex++ {
		startsAt := eachPageItems * pageIndex
		endAt := startsAt + eachPageItems

		go GetAsyncWithRange(
			&processedItems,
			content,
			allRawItems,
			genericProcessor,
			startsAt,
			endAt,
			&wg)
	}

	wg.Wait()

	return &processedItems
}
