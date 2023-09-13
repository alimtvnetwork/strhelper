package sprocess

import (
	"gitlab.com/auk-go/strhelper/strhelpercore"
)

func Get(
	content *interface{},
	allRawItems []interface{},
	genericProcessor strhelpercore.GenericProcessor,
) []interface{} {
	processedItems := make([]interface{}, len(allRawItems))

	args := strhelpercore.GenericProcessorArgs{
		Content:              content,
		RawSplitContents:     allRawItems,
		Index:                -1,
		SingleContentAtIndex: nil,
	}

	for args.Index, args.SingleContentAtIndex = range allRawItems {
		processedItems[args.Index] = genericProcessor(&args)
	}

	return processedItems
}
