package pk2

import (
	"encoding/json"
	"fmt"
	
	"gitlab.com/auk-go/core/coredata/corejson"
	"gitlab.com/auk-go/core/iserror"
)

type CustomField struct {
	SomeField int
}

type JsonExampleOutput struct {
	FF1         string `json:",omitempty"`
	FAlim       string `json:",omitempty"`
	customField CustomField
}

type JsonExample struct {
	F1          string `json:",omitempty"`
	Alim        string `json:",omitempty"`
	customField CustomField
}

func (it *JsonExample) String() string {
	return "x"
}

func (it JsonExample) MarshalJSON() ([]byte, error) {
	it.Alim = "JSON Marshalling is About to happen"
	fmt.Println("invoking custom - marshal")
	
	return json.Marshal(JsonExampleOutput{
		FF1:   it.F1,
		FAlim: it.Alim,
		customField: CustomField{
			SomeField: 12,
		},
	})
}

func (it *JsonExample) UnmarshalJSON(data []byte) error {
	it.Alim = "JSON Marshalling is About to happen"
	fmt.Println("invoking custom - unmarshall")
	var x JsonExampleOutput
	err := json.Unmarshal(data, &x)
	
	if iserror.Empty(err) {
		it.Alim = x.FAlim
		it.F1 = x.FF1
		it.customField = x.customField
	}
	
	return err
}

func (it *JsonExample) AsJsonMarshaller() corejson.JsonMarshaller {
	return it
}
