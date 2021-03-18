package indextests

import (
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/corecompare"

	"gitlab.com/evatix-go/core/coretests"

	"gitlab.com/evatix-go/strhelper/index"
	"gitlab.com/evatix-go/strhelper/tests/testwrappers/indextestwrappers"
)

func TestOfAllPtr(t *testing.T) {
	for i, testCase := range indextestwrappers.OfAllPtrCasesPtr {
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
		actual := index.OfAllPtr(
			&testCase.Content,
			&testCase.SearchingContent,
			testCase.InitializedPosition,
			testCase.Limits,
			testCase.IsCaseSensitive,
		)

		testCase.SetActual(actual)

		Convey(testHeader, t, func() {
			// Assert
			Convey(coretests.GetAssertMessage(caseMessenger, i), func() {
				isSame := corecompare.IntArrayPtr(actual, expected)

				if !isSame {
					header := "\n ==================Actual vs Expectation==================\nExpectations : "
					log.Println(header, expected)
					log.Println("Actual : ", actual)
				}

				So(isSame, ShouldBeTrue)
			})
		})
	}
}
