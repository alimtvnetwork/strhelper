package lines

import (
	"sync"

	"gitlab.com/evatix-go/strhelper/strconst"
	"gitlab.com/evatix-go/strhelper/strhelpercore"
)

func Process(
	content *string,
	lines *[]string,
	lineProcessor *strhelpercore.LineProcessor,
) *[]*string {
	newStrings := make([]*string, len(*lines))

	args := strhelpercore.LineArgs{
		Content: content,
		Lines:   lines,
		Index:   -1,                   // it will change per line
		Line:    strconst.EmptyString, // it will change per line
	}

	for args.Index, args.Line = range *lines {
		newStrings[args.Index] = (*lineProcessor)(&args)
	}

	return &newStrings
}

func parallelProcessFunc(
	processedStrings *[]*string,
	lineProcessor *strhelpercore.LineProcessor,
	args *strhelpercore.LineArgs,
	wg *sync.WaitGroup,
) {
	// example : https://bit.ly/3lLndEF
	defer wg.Done()
	(*processedStrings)[(*args).Index] = (*lineProcessor)(args)
}

// Runs loop in async mode (in golang starts with go).
//
// It requires more memory to deal with parallel execution.
//
// Sometimes it is faster for large collection in async, memory is cheap than idle cpu.
func ProcessAsync(
	content *string,
	lines *[]string,
	lineProcessor *strhelpercore.LineProcessor,
) *[]*string {
	length := len(*lines)
	newStrings := make([]*string, length)
	var wg sync.WaitGroup
	wg.Add(length)

	for index, line := range *lines {
		args := strhelpercore.LineArgs{
			Content: content,
			Lines:   lines,
			Index:   index, // it will change per line
			Line:    line,  // it will change per line
		}

		go parallelProcessFunc(
			&newStrings,
			lineProcessor,
			&args,
			&wg)
	}

	wg.Wait()

	return &newStrings
}
