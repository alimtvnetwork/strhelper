package paddingbenchmark

import (
	"strings"
	"testing"

	"gitlab.com/evatix-go/strhelper/padding"
	"gitlab.com/evatix-go/strhelper/tests/testwrappers/paddingwrappers"
)

func BenchmarkStringsRepeat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, benchmarkCase := range paddingwrappers.RepeatTestCases {
			strings.Repeat(
				benchmarkCase.Content,
				benchmarkCase.RepeatWidth)
		}
	}
}

func BenchmarkPaddingRepeat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, benchmarkCase := range paddingwrappers.RepeatTestCases {
			padding.Repeat(
				benchmarkCase.Content,
				benchmarkCase.RepeatWidth)
		}
	}
}
