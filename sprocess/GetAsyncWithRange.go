package sprocess

import (
	"sync"

	"gitlab.com/evatix-go/strhelper/internal/panichelper"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

// Runs loop in async mode (in golang starts with go).
//
// It requires more memory to deal with parallel execution.
//
// Sometimes it is faster for large collection in async, memory is cheap than idle cpu.
func GetAsyncWithRange(
	processedItems *[]*interface{},
	content *interface{},
	allRawItems *[]interface{},
	genericProcessor *strhelpercore.GenericProcessor,
	startAtIndex int,
	endAtIndex int,
	parentWg *sync.WaitGroup,
) *[]*interface{} {
	length := endAtIndex - startAtIndex
	defer parentWg.Done()

	if length < 0 || startAtIndex < 0 {
		panichelper.StartAtIndexFailed(startAtIndex, length)
	}

	totalLength := len(*allRawItems)

	if endAtIndex > totalLength-1 {
		endAtIndex = totalLength - 1
	}

	var wg sync.WaitGroup
	wg.Add(length)

	for ; startAtIndex <= endAtIndex; startAtIndex++ {
		copiedContentAtIndex := (*allRawItems)[startAtIndex]
		args := strhelpercore.GenericProcessorArgs{
			Content:              content,
			RawSplitContents:     allRawItems,
			Index:                startAtIndex,
			SingleContentAtIndex: &copiedContentAtIndex,
		}

		go parallelProcessFunc(
			processedItems,
			genericProcessor,
			&args,
			&wg)
	}

	wg.Wait()

	return processedItems
}
