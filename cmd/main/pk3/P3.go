package pk3

import (
	"fmt"
	"strconv"
	
	"gitlab.com/auk-go/strhelper/cmd/main/pk1"
	"gitlab.com/auk-go/strhelper/cmd/main/pk2"
)

type P3 struct {
}

func (receiver P3) Do1(x int) int {
	fmt.Println(strconv.Itoa(x) + "-p3")
	
	return pk1.P1{}.Do1(x)
}

func (receiver P3) Do2(x int) {
	pk2.P2{}.Do2(x)
	
	fmt.Println(strconv.Itoa(x) + "-p3")
}

func (receiver P3) Do3(x int) {
	fmt.Println(strconv.Itoa(x) + "-p3")
}
