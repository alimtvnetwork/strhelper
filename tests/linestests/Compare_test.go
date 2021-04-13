package linestests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/strhelper/strlines"
	"gitlab.com/evatix-go/strhelper/tests/testwrappers/linestestwrappers"
)

func Test_Compare(t *testing.T) {
	for i, testCase := range linestestwrappers.CompareTestCases {
		// Arrange
		testHeader := coretests.GetTestHeader(testCase)

		// Act
		actual := strlines.Compare(
			testCase.LeftLines,
			testCase.RightLines,
			testCase.StartsAt,
			testCase.IsPanicOnLengthDifferent,
			testCase.IsCaseSensitive)

		testCase.SetActual(actual)

		// Assert
		Convey(testHeader, t, func() {
			Convey(coretests.GetAssertMessage(testCase, i), func() {
				So(actual, ShouldEqual, testCase.Expected())
			})
		})
	}
}
