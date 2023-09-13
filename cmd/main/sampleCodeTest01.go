package main

import (
	"bytes"
	"fmt"
	"strings"

	"gitlab.com/auk-go/core/coredata/stringslice"
	"gitlab.com/auk-go/strhelper/anyto"
	"gitlab.com/auk-go/strhelper/strlines"
	"gitlab.com/auk-go/strhelper/strs/isstrs"
	"gitlab.com/auk-go/strhelper/whitespace"
)

func sampleCodeTest01() {
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

	leftBytes := strlines.ToUnsafeBytesPtr(&leftUpto3)
	rightBytes := strlines.ToUnsafeBytesPtr(&rightLines)
	comparedResult5 := isstrs.BytesEquals(*leftBytes, *rightBytes, 0)

	fmt.Println(comparedResult5)

	leftBytes2, _ := anyto.Bytes(leftUpto3)
	rightBytes2, _ := anyto.Bytes(rightLines)
	comparedResult6 := bytes.Equal(leftBytes2, rightBytes2)

	fmt.Println(comparedResult6)

	left2Lines := []string{
		"Line 1",
		"Line 2",
		"Line 3",
	}

	leftBytes3, _ := anyto.Bytes(left2Lines)

	comparedResult7 := bytes.Equal(leftBytes3, *rightBytes)
	fmt.Println(comparedResult7)

	whitespaceTest := "testing o  \t \n\n\n "
	fmt.Println((*whitespace.AllWhitespacesRuneIndexesMap(&whitespaceTest))['\t'])
}
