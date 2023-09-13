package indextests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"

	"gitlab.com/auk-go/strhelper/stringindex"
	"gitlab.com/auk-go/strhelper/tests/testwrappers/indextestwrappers"
)

func Test_OfLastAllPtr(t *testing.T) {
	for i, testCase := range indextestwrappers.OfLastAllPtrCasesPtr {
		// Validate
		if testCase.HasPanic {
			// will not work with panic cases
			continue
		}

		// Arrange
		caseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(
			caseMessenger)
		expected := testCase.ExpectedAsIntArray()

		// Act
		actual := stringindex.OfLastAllPtr(
			testCase.Content,
			testCase.SearchingContent,
			testCase.InitializedPosition,
			testCase.Limits,
			testCase.IsCaseSensitive,
		)

		testCase.SetActual(actual)

		// Assert
		convey.Convey(testHeader, t, func() {
			convey.Convey(coretests.GetAssertMessage(caseMessenger, i), func() {
				convey.So(actual, convey.ShouldResemble, expected)
			})
		})
	}
}
