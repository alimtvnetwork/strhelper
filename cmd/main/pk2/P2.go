package pk2

import (
	"fmt"
	"strconv"
	
	"gitlab.com/auk-go/strhelper/cmd/main/pk1"
)

type P2 struct {
}

func (receiver P2) Do2(x int) {
	fmt.Println(strconv.Itoa(x) + "some")
}

func (receiver P2) MyCustomImplementation(r pk1.Requester) pk1.Outputter {
	// future
	rField := r.Field1()
	output := Pk2Output{}
	output.SetResult("doing something in future from Pk2" + rField)
	
	return &output
}
