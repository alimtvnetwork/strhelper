package testscore

import (
	"fmt"
)

func GetTestHeader(testCaseMessager TestCaseMessager) string {
	return fmt.Sprintf("Method : [%s]",
		testCaseMessager.FuncName(),
	)
}
