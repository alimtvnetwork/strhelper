package main

import (
	"fmt"

	"gitlab.com/auk-go/strhelper/attrmeta"
	"gitlab.com/auk-go/strhelper/cmd/main/pk1"
	"gitlab.com/auk-go/strhelper/cmd/main/pk2"
	"gitlab.com/auk-go/strhelper/splits"
)

type SomeInterface interface {
	SomeMethodWithY() string
	SomeNewInterfaceMethod() int
}

type Field1Encapsulater interface {
	Field1() string
	SetField1(x string)
}

type SomeStruct struct {
	A, B   float64 // Received Answer : nil
	field1 string
}

// Field1
//
//	Getter
func (it *SomeStruct) Field1() string {
	return it.field1
}

// SetField1
//
//	Setter
func (it *SomeStruct) SetField1(field1 string) {
	it.field1 = field1
}

type SomeStruct2 struct {
	A, B string
}

type SomeStructY struct {
	A, B string
}

func (receiver SomeStructY) SomeMethodWithY() string {
	return "something from Y"
}

type SomeCombinedStruct struct {
	// SomeStructY
	SomeStruct2
	SomeStruct
	A int
}

func RealWorldFunc(encapsulater Field1Encapsulater) string {
	x := encapsulater.Field1()

	fmt.Println(x)

	return x
}

func NewPtr() *SomeStruct {
	n1 := SomeStruct{}

	return &n1
}

func main() {
	// var x SomeCombinedStruct
	// m := x.SomeStructY // in terms of memory (no copy), access slow if pointer

	// fmt.Println(RealWorldFunc(NewPtr()))

	x := pk2.Hello2{}
	hello1 := pk1.Hello1{}

	fmt.Println(x.Do(&hello1))

	// fmt.Println("hello World")
	// sampleCodeTest01()
	// splitsTest01()
	// collectionTest01()
}

func collectionTest01() {
	collection := attrmeta.NewCollectionDefault()

	collection.Str("something", "some data").Strings(
		"slice",
		"some val",
		"soime val2",
	)

	// fmt.Println(collection.StackTracesDefault())
	collection.LogWithTraces()
	// collection.LogWithTraces()
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
