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
