package pk2

import "gitlab.com/auk-go/strhelper/cmd/main/pk1"

type Pk2Output struct {
	result string
}

func (it Pk2Output) Result() string {
	return it.result
}

func (it *Pk2Output) SetResult(x string) {
	// TODO implement me
	it.result = x
}

func (it Pk2Output) AsOutputter() pk1.Outputter {
	return &it
}
