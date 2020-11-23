package tests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"gitlab.com/evatix-go/strhelper/isstr"
	"gitlab.com/evatix-go/strhelper/tests/testscore"
	"gitlab.com/evatix-go/strhelper/tests/testwrappers"
)

func TestIsEndsWith(t *testing.T) {
	for i, testCase := range testwrappers.EndsWithTestCases {
		// Arrange
		testHeader := testscore.GetTestHeader(testCase)

		// Act
		actual := isstr.EndsWith(
			testCase.WholeText,
			testCase.Search,
			testCase.StartsAt,
			testCase.IsCaseSensitive)

		testCase.SetActual(actual)

		Convey(testHeader, t, func() {
			// Assert
			Convey(testscore.GetAssertMessage(testCase, i), func() {
				So(actual, ShouldEqual, testCase.Expected())
			})
		})
	}
}
