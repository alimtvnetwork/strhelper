package main

import (
	"errors"
	"fmt"

	"gitlab.com/auk-go/errorwrapper/errnew"
	"gitlab.com/auk-go/errorwrapper/errtype"
	"gitlab.com/auk-go/errorwrapper/errwrappers"
	"gitlab.com/auk-go/errorwrapper/ref"
	"gitlab.com/auk-go/strhelper/attrmeta"
	"gitlab.com/auk-go/strhelper/splits"
)

func main() {
	// fmt.Println("hello World")
	// sampleCodeTest01()
	// splitsTest01()
	// collectionTest01()

	SomeMethod()

	// errWrap1 := errnew.Null.Error(errors.New("something nil"))
	// errWrap2 := errnew.Error.TypeOnly(errtype.DomainMissing)
	//
	// finalErrWrap := errWrap1.ConcatNew().Wrapper(errWrap2)
	//
	// fmt.Println(finalErrWrap.FullStringWithTraces())
	// list := errwrappers.Empty()
	//
	// list.Add(errtype.NotFoundGroup)
	// list.AddError(errors.New("some error"))
	// list.AddWrapperPtr(finalErrWrap)
	// list.AddTypeRefQuick(errtype.NotFoundName, "rahim", "zakaria")
	// // list.AddUsingMsg(errtype.NotFoundGroup, "group name is not found")
	// // list.AddRef1Msg(errtype.Directory, "directory not found", "xy", "path/path2")
	// // list.AddPathIssue(errtype.PathMissingOrInvalid, errors.New("some path issue"), "path/path2")
	//
	// fmt.Println(list.FullStringsWithTraces())
}

func SomeMethod() {
	x := errors.New("something error")
	errWrap := errnew.Path.File(x, "/path/path2")
	errWrap2 := errnew.Path.File(x, "/pathX/path2")

	fmt.Println(errWrap.FullStringWithTraces())
	errWrap.ConcatNew().Wrapper(errWrap2)
	f1 := errnew.Ref.Many(errtype.DuplicateIssue, ref.Value{
		Variable: "x",
		Value:    []int{1, 5, 6},
	}, ref.Value{
		Variable: "y",
		Value:    []string{"alim", "alim 2"},
	})

	fmt.Println(f1.FullStringWithTraces())
	// if x != nil {
	// 	y := errors.New("y")
	//
	// 	if y != nil {
	// 		finalErr := errors.New(x.Error() + y.Error())
	//
	// 		fmt.Println(finalErr)
	// 	}
	//
	// }

	list := errwrappers.Empty()

	list.AddRef(errtype.IdConflict, errWrap2.Error(), "path", []string{"1x", "2x"})

	fmt.Println(list.FullStringsWithTraces())
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
