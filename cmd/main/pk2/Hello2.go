package pk2

import "gitlab.com/auk-go/strhelper/cmd/main/pk1"

type Hello2 struct {
}

func (it *Hello2) Do(hello1 *pk1.Hello1) string {
	return "Hello 2 - depends on - " + hello1.Do()
}
