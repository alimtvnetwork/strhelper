package main

import (
	"fmt"
	"unicode"

	"gitlab.com/evatix-go/strhelper"
	"gitlab.com/evatix-go/strhelper/charhelper"
	"gitlab.com/evatix-go/strhelper/concat"
	"gitlab.com/evatix-go/strhelper/whitespace"
)

func main() {
	a := "wcaabdwdwdwwdcabcabbccacacccadecayucabwawwweedwwcabaa"
	b := "cab"
	//
	// fmt.Println(strhelper.IndexesOfAllPtr(&a, &b, 0, false)) // [13 16 35 48] language default
	// fmt.Println(strhelper.IndexesOfAllPtr(&a, &b, 1, true))  // [13 16 35 48]

	fmt.Println(strhelper.LastIndexOfPtr(&a, &b, 0, true))  // 48 language default
	fmt.Println(strhelper.LastIndexOfPtr(&a, &b, 1, false)) // 48

	a2 := "xcabhellophwihwcab..CAB"
	replacedWith := "-REPLACED ALIM-"

	fmt.Println(strhelper.LongestCommonSuffixCountPtr(&a2, &b, 0, false)) // expects 3
	fmt.Println(strhelper.LongestCommonSuffixCountPtr(&a, &b, 0, false))  // expects 0
	fmt.Println(strhelper.LongestCommonSuffixCountPtr(&a2, &b, 1, false)) // expects 2
	fmt.Println(strhelper.LongestCommonSuffixCountPtr(&a2, &b, 2, false)) // expects 1
	fmt.Println(strhelper.LongestCommonSuffixCountPtr(&a2, &b, 3, false)) // expects 0
	// fmt.Println(strhelper.ReplacePtr(&a2, &b, &replacedWith, 0,-1, true)) // expects "x-REPLACED ALIM-hellophwihw-REPLACED ALIM-"
	// fmt.Println(strhelper.ReplacePtr(&a2, &b, &replacedWith, 0,-1, false)) // expects "x-REPLACED ALIM-hellophwihw-REPLACED ALIM"
	// fmt.Println(strhelper.ReplacePtr(&a2, &b, &replacedWith, 1,1, false)) // expects "x-REPLACED ALIM-hellophwihwcab"
	// fmt.Println(strhelper.ReplacePtr(&a2, &b, &replacedWith, 2,2, false)) // expects "x-REPLACED ALIM-hellophwihwcab"
	fmt.Println(strhelper.ReplacePtr(&a2, &b, &replacedWith, 2, 2, true)) // expects "xcabhellophwihw-REPLACED ALIM-..CAB"

	// replaceMap := map[string]string{
	// 	"cab" : "-cabReplacer       V1-",
	// 	"wih" : "-wihReplacer       V1-",
	// 	"hello" : "-helloReplacer       V1-",
	// }
	//
	// fmt.Println(strhelper.ReplaceMultiple(a2, replaceMap, 0,-1, false)) // expects "xcabhellophwihw-REPLACED ALIM-..CAB"
	// replaceMap[""] = "Hello"
	// fmt.Println(strhelper.ReplaceMultiple("", replaceMap, 0,-1, false)) // expects "xcabhellophwihw-REPLACED ALIM-..CAB"
	asciiWhitespacesArray := whitespace.GetAscIISpaceArray()
	asciiWhitespacesArray['p'] = 1
	asciiWhitespacesArray2 := whitespace.GetAscIISpaceArray()
	fmt.Println((asciiWhitespacesArray2))
	fmt.Println(charhelper.AscIIArrayToString(&asciiWhitespacesArray))
	fmt.Println(strhelper.RemoveWhitespaces("                                                   Hello No Sapce               Alim                                  ", 0)) // expects "HelloNoSapceAlim"
	fmt.Println(strhelper.RemoveWhitespaces("                                                   Hello No Sapce               Alim                                  ", 5)) // expects "     HelloNoSapceAlim"
	fmt.Println(strhelper.RemoveNewLines("\r\n\f\n\r\f Hello\fNo\fSapce\fAlim\f", 0))                                                                                     // expects " HelloNoSapceAlim"
	fmt.Println(strhelper.RemoveNewLines("\r\n\f\n\r\f Hello\fNo\fSapce\fAlim\f", 2))                                                                                     // expects "\r\n HelloNoSapceAlim"
	fmt.Println(strhelper.Reverse("Hello World"))
	fmt.Println(strhelper.Reverse("Hello Worldx"))
	fmt.Println(strhelper.Reverse(""))
	fmt.Println(charhelper.CountAscIIChars("   Hellop", &asciiWhitespacesArray, 0, true))
	fmt.Println(unicode.MaxLatin1) // 255
	fmt.Println(rune(("世界")[0]))   // 228
	fmt.Println([]rune("世界")[0])   // 19990
	fmt.Println(0x85)
	strToReplace := "hello,-doyou/trymeNO--"
	splitResults := strhelper.MultipleQuickSplits(strToReplace, ",", "-", "--")
	fmt.Println(concat.AnyFullNameValues("", false, *splitResults))
}
