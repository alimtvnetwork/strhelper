package sprocess

import (
	"sync"

	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// GetAsync Runs loop in async mode (in golang starts with go).
//
// It requires more memory to deal with parallel execution.
//
// Sometimes it is faster for large collection in async, memory is cheap than idle cpu.
func GetAsync(
	content interface{},
	allRawItems []interface{},
	genericProcessor strhelpercore.GenericProcessor,
) []interface{} {
	length := len(allRawItems)
	processedItems := make([]interface{}, length)
	var wg sync.WaitGroup
	wg.Add(length)

	for index, contentAtIndex := range allRawItems {
		copiedContentAtIndex := contentAtIndex
		args := strhelpercore.GenericProcessorArgs{
			Content:              content,
			RawSplitContents:     allRawItems,
			Index:                index,
			SingleContentAtIndex: &copiedContentAtIndex,
		}

		go parallelProcessFunc(
			processedItems,
			genericProcessor,
			args,
			&wg)
	}

	wg.Wait()

	return processedItems
}
