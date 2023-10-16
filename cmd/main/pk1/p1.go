package pk1

import (
	"fmt"
	"strconv"
)

type P1 struct {
}

func (receiver P1) Do1(x int) int {
	fmt.Println(strconv.Itoa(x))
	
	return x + 5
}

func (receiver P1) DoSomethingInFuture(
	processor Processor1,
	r Requester,
) Outputter {
	output := processor(r) // future
	
	// modify
	output.SetResult(output.Result() + "-some changes from pk1 -")
	
	return output
}
