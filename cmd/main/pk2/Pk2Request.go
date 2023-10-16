package pk2

import "gitlab.com/auk-go/strhelper/cmd/main/pk1"

type Pk2Request struct {
	field1 string
}

func (it *Pk2Request) SetField1(s string) {
	// TODO implement me
	it.field1 = s
}

func (it Pk2Request) Field1() string {
	return it.field1
}

func (it Pk2Request) AsRequester() pk1.Requester {
	return &it
}
