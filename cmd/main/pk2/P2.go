package pk2

import (
	"fmt"
	"strconv"
)

type P2 struct {
}

func (receiver P2) Do2(x int) {
	fmt.Println(strconv.Itoa(x) + "some")
}
