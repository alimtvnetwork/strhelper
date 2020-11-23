package linestests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/strhelper/lines"
	"gitlab.com/evatix-go/strhelper/tests/testscore"
	"gitlab.com/evatix-go/strhelper/tests/testwrappers/linestestwrappers"
)

func TestCompare(t *testing.T) {
	for i, testCase := range linestestwrappers.CompareTestCases {
		// Arrange
		testHeader := testscore.GetTestHeader(testCase)

		// Act
		actual := lines.Compare(
			testCase.LeftLines,
			testCase.RightLines,
			testCase.StartsAt,
			testCase.IsPanicOnLengthDifferent,
			testCase.IsCaseSensitive)

		testCase.SetActual(actual)

		// Assert
		Convey(testHeader, t, func() {
			Convey(testscore.GetAssertMessage(testCase, i), func() {
				So(actual, ShouldEqual, testCase.Expected())
			})
		})
	}
}
