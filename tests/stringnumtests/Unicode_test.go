package stringnumtests

import (
	"testing"
)

func Test_Unicode_IntegerSigned(t *testing.T) {
	numberAssertion(t, unicodeIntegerSignedTestCases)
}

func Test_Unicode_IntegerUnsigned(t *testing.T) {
	numberAssertion(t, unicodeIntegerUnsignedTestCases)
}

func Test_Unicode_IntegerUnsignedWithoutPlus(t *testing.T) {
	numberAssertion(t, unicodeIntegerUnsignedWithoutPlusTestCases)
}

func Test_Unicode_NumberSigned(t *testing.T) {
	// represents float numbers signed
	numberAssertion(t, unicodeNumberSingedTestCases)
}

func Test_Unicode_NumberUnsigned(t *testing.T) {
	// represents float numbers unsigned, including plus + symbol
	numberAssertion(t, unicodeNumberUnsignedTestCases)
}

func Test_Unicode_NumberUnsignedWithoutPlus(t *testing.T) {
	// represents float numbers unsigned without plus symbol
	numberAssertion(t, unicodeNumberUnsignedWithoutPlusTestCases)
}
