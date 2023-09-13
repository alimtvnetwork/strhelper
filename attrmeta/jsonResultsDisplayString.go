package attrmeta

import (
	"fmt"

	"gitlab.com/auk-go/core/constants"
	"gitlab.com/auk-go/core/corecsv"
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/coreinterface/serializerinf"
	"gitlab.com/auk-go/core/errcore"
)

func jsonResultsDisplayString(jsonResults ...*corejson.Result) string {
	if len(jsonResults) == 0 {
		return ""
	}

	slice := make([]string, len(jsonResults))

	for i, result := range jsonResults {
		slice[i] = jsonResultToDisplayString(result)
	}

	return corecsv.DefaultCsv(slice...)
}

func jsonResultToDisplayString(
	jsonResult serializerinf.BaseJsonResulter,
) string {
	if jsonResult == nil {
		return constants.NilAngelBracket
	}

	if jsonResult.HasError() {
		return fmt.Sprintf(
			"Value : %s, Error: %s, Type: %s",
			jsonResult.SafeString(),
			errcore.ToString(jsonResult.MeaningfulError()),
			jsonResult.BytesTypeName())
	}

	return string(jsonResult.SafeValues())
}
