package main

import (
	"fmt"
	"strings"

	"gitlab.com/evatix-go/core/coredata/stringslice"
	"gitlab.com/evatix-go/strhelper/whitespace"

	"gitlab.com/evatix-go/strhelper/anyto"
	"gitlab.com/evatix-go/strhelper/strlines"
	"gitlab.com/evatix-go/strhelper/strs/isstrs"
)

func main() {
	// fmt.Println("hello World")
	leftLines := []string{
		"Line 1",
		"Line 2",
		"Line 3",
		"Line 4",
		"Line 5",
	}

	rightLines := []string{
		"Line 1",
		"Line 2",
		"Line 3",
	}

	slice := stringslice.MergeNew(leftLines, rightLines...)

	fmt.Println(slice)

	slice2 := stringslice.PrependNew(leftLines, rightLines...)

	fmt.Println(slice2)

	comparedResult := strlines.Compare(
		&leftLines,
		&rightLines,
		0, false, true)

	fmt.Println(comparedResult)
	fmt.Println(strings.Compare("a", ""))

	comparedResult2 := strlines.Compare(nil, &rightLines, 0, false, true)

	fmt.Println(comparedResult2)
	fmt.Println(strings.Compare("", "a"))

	comparedResult3 := strlines.Compare(&leftLines, nil, 0, false, true)

	fmt.Println(comparedResult3)
	fmt.Println(strings.Compare("a", ""))

	leftUpto3 := leftLines[0:3]
	comparedResult4 := strlines.Compare(&leftUpto3, &rightLines, 1, false, true)

	fmt.Println(comparedResult4)
	fmt.Println(strings.Compare("a", "a"))

	leftBytes := strlines.ToUnsafeBytes(&leftUpto3)
	rightBytes := strlines.ToUnsafeBytes(&rightLines)
	comparedResult5 := isstrs.BytesEquals(leftBytes, rightBytes, 0)

	fmt.Println(comparedResult5)

	leftBytes, _ = anyto.Bytes(leftUpto3)
	rightBytes, _ = anyto.Bytes(rightLines)
	comparedResult6 := isstrs.BytesEquals(leftBytes, rightBytes, 0)

	fmt.Println(comparedResult6)

	left2Lines := []string{
		"Line 1",
		"Line 2",
		"Line 3",
	}

	leftBytes2, _ := anyto.Bytes(left2Lines)

	comparedResult7 := isstrs.BytesEquals(leftBytes2, rightBytes, 0)
	fmt.Println(comparedResult7)

	whitespaceTest := "testing o  \t \n\n\n "
	fmt.Println((*whitespace.AllWhitespacesRuneIndexesMap(&whitespaceTest))['\t'])
}
