package stringnumtests

import (
	"testing"
)

func Test_AscII_IntegerSigned(t *testing.T) {
	numberAssertion(t, asciiIntegerSignedTestCases)
}

func Test_AscII_IntegerUnsigned(t *testing.T) {
	numberAssertion(t, asciiIntegerUnsignedTestCases)
}

func Test_AscII_IntegerUnsignedWithoutPlus(t *testing.T) {
	numberAssertion(t, asciiIntegerUnsignedWithoutPlusTestCases)
}

func Test_AscII_NumberSigned(t *testing.T) {
	// represents float numbers signed
	numberAssertion(t, asciiNumberSignedTestCases)
}

func Test_AscII_NumberUnsigned(t *testing.T) {
	// represents float numbers unsigned including plus symbol +
	numberAssertion(t, asciiNumberUnsignedTestCases)
}

func Test_AscII_NumberUnsignedWithoutPlus(t *testing.T) {
	// represents float numbers unsigned without plus symbol
	numberAssertion(t, asciiNumberUnsignedWithoutPlusTestCases)
}
