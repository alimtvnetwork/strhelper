package splitstests

import (
	"log"
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/strhelper/internal/isstrsinternal"
	"gitlab.com/evatix-go/strhelper/splits"
	"gitlab.com/evatix-go/strhelper/tests/testwrappers/splitstestwrapper"
)

func Test_LastByRune(t *testing.T) {
	for i, testCase := range splitstestwrapper.LastByRuneTestCases {
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
		actual := splits.LastByRunePtr(
			testCase.Content,
			testCase.SearchingContent,
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
