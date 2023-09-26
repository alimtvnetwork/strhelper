package pk1

import "gitlab.com/auk-go/strhelper/cmd/main/pk2"

type Hello1 struct {
}

func (it *Hello1) Do() string {
	return "Hello 1"
}

func (it *Hello1) DoFromPk2(hello2 pk2.Hello2) string {
	return "Hello 1 - depends on -" + hello2.Do(it)
}
