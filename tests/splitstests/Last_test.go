package splitstests

import (
	"log"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/auk-go/core/coretests"

	"gitlab.com/auk-go/strhelper/internal/isstrsinternal"
	"gitlab.com/auk-go/strhelper/splits"
	"gitlab.com/auk-go/strhelper/tests/testwrappers/splitstestwrapper"
)

func Test_Last(t *testing.T) {
	for i, testCase := range splitstestwrapper.LastTestCases {
		// Validate
		if testCase.HasPanic {
			// will not work with panic cases
			continue
		}

		// Arrange
		caseMessenger := testCase.AsTestCaseMessenger()
		testHeader := coretests.GetTestHeader(
			caseMessenger)
		expected := testCase.ExpectedAsStringsArray()

		// Act
		actual := splits.LastByLimit(
			testCase.Content,
			testCase.SearchingContent,
			testCase.IsCaseSensitive,
			testCase.Limits,
		)

		testCase.SetActual(actual)

		// Assert
		convey.Convey(testHeader, t, func() {
			convey.Convey(coretests.GetAssertMessage(caseMessenger, i), func() {
				isSame := isstrsinternal.Equals(actual, expected, 0, true)

				if !isSame {
					header := "\n ==================Actual vs Expectation==================\nExpectations : "
					log.Println(header, expected)
					log.Println("Actual : ", actual)
				}

				convey.So(isSame, convey.ShouldBeTrue)
			})
		})
	}
}
