package sprocess

import (
	"sync"

	"gitlab.com/auk-go/strhelper/strhelpercore"
)

func parallelProcessFunc(
	processedItems []interface{},
	genericProcessor strhelpercore.GenericProcessor,
	args strhelpercore.GenericProcessorArgs,
	wg *sync.WaitGroup,
) {
	// example : https://bit.ly/3lLndEF
	defer wg.Done()
	processedItems[args.Index] = genericProcessor(&args)
}
