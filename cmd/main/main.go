package main

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper"
)

func main() {
	a := "wcaabdwdwdwwdcabcabbccacacccadecayucabwawwweedwwcabaa"
	b := "cab"

	fmt.Println(strhelper.IndexesOfAllPtr(&a, &b, 0, false)) // [13 16 35 48] language default
	fmt.Println(strhelper.IndexesOfAllPtr(&a, &b, 1, true))  // [13 16 35 48]

	fmt.Println(strhelper.LastIndexOfPtr(&a, &b, 0, true))   // 48 language default
	fmt.Println(strhelper.LastIndexOfPtr(&a, &b, 1, false))  // 48
}
