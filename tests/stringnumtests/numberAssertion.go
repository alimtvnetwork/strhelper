package stringnumtests

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
	"gitlab.com/evatix-go/core/conditional"
	"gitlab.com/evatix-go/core/errcore"
)

func numberAssertion(
	t *testing.T,
	numberVerifyGroupWrapper *numberVerifyGroupWrapper,
) {
	testerFunc := numberVerifyGroupWrapper.verifierFunc
	methodNameWrapped := numberVerifyGroupWrapper.MethodName()

	for _, wrapper := range numberVerifyGroupWrapper.testCases {
		isValidNumber := testerFunc(wrapper.input)

		numberString := conditional.String(
			wrapper.isValidNumber,
			"be a number\n",
			"NOT be a number\n")

		header := errcore.ExpectingSimpleNoType(
			methodNameWrapped+wrapper.input+" should be "+numberString,
			wrapper.isValidNumber,
			isValidNumber)

		isValid := isValidNumber == wrapper.isValidNumber

		convey.Convey(header, t, func() {
			errcore.FailedPrint(!isValid, header)

			convey.So(isValid, convey.ShouldBeTrue)
		})
	}
}
