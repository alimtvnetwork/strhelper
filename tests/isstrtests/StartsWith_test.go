package isstrtests

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"

	"gitlab.com/auk-go/strhelper/isstr"
	"gitlab.com/auk-go/strhelper/tests/testwrappers"
)

func TestStartsWith(t *testing.T) {
	for i, testCase := range testwrappers.StartsWithTestCases {
		// Arrange
		testHeader := coretests.GetTestHeader(testCase)

		// Act
		actual := isstr.StartsWith(
			testCase.WholeText,
			testCase.Search,
			testCase.StartsAt,
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
