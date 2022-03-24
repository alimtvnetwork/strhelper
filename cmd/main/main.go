package main

import (
	"fmt"

	"gitlab.com/evatix-go/strhelper/attrmeta"
	"gitlab.com/evatix-go/strhelper/splits"
)

func main() {
	// fmt.Println("hello World")
	// sampleCodeTest01()
	// splitsTest01()
	collectionTest01()
}

func collectionTest01() {
	collection := attrmeta.NewCollectionDefault()

	collection.Str("something", "some data").Strings(
		"slice",
		"some val",
		"soime val2",
	)

	collection.StackTracesDefault()
	collection.LogWithTraces()
}

func splitsTest01() {
	slice := []string{
		"some=val",
		"some1=val2",
		"some4=val3",
	}

	keyvals := splits.ByEqual(slice[0])

	fmt.Println(keyvals)
}
