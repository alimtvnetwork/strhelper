package paddingtests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/strhelper/padding"
	"gitlab.com/evatix-go/strhelper/tests/testwrappers/paddingwrappers"
)

func Test_RepeatPtr(t *testing.T) {
	for i, testCase := range paddingwrappers.RepeatTestCases {
		// Validate
		if testCase.HasPanic {
			// will not work with panic cases
			continue
		}

		// Arrange
		caseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(
			caseMessenger)
		expected := testCase.Expected()

		// Act
		actual := padding.RepeatPtr(
			&testCase.Content,
			testCase.RepeatWidth,
		)

		testCase.SetActual(actual)

		// Assert
		convey.Convey(testHeader, t, func() {
			convey.Convey(coretests.GetAssertMessage(caseMessenger, i), func() {
				convey.So(actual, convey.ShouldEqual, expected)
			})
		})
	}
}
