![CIMUX golang strhelper logo](https://gitlab.com/evatix-go/strhelper/uploads/12f1449174f417be0570627cf6e749ea/go-string-250.png)

# Strings Extension Introduction (`strhelper`)

Go Strings library additional methods, simplification of string modification and verification.

## Git Clone

`git clone https://gitlab.com/evatix-go/strhelper.git`

### 2FA enabled, for linux

`git clone https://[YourGitLabUserName]:[YourGitlabAcessTokenGenerateFromGitlabsTokens]@gitlab.com/evatix-go/strhelper.git`

### Prerequisites

- Update git to latest 2.29
- Update or install the latest of Go 1.15.2
- Either add your ssh key to your gitlab account
- Or, use your access token to clone it.

## Installation

`go get gitlab.com/evatix-go/strhelper`

### Go get issue for private package

- Update git to 2.29
- Enable go modules. (Windows : `go env -w GO111MODULE=on`, Unix : `export GO111MODULE=on`)
- Add `gitlab.com/evatix-go` to go env private

To set for Windows:

`go env -w GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

To set for Unix:

`expoort GOPRIVATE=[AddExistingOnes;]gitlab.com/evatix-go`

## Build Fix

```cmd
go get gitlab.com/evatix-go/core
go get gitlab.com/evatix-go/errorwrapper

go mod download gitlab.com/evatix-go/core
go mod download gitlab.com/evatix-go/errorwrapper
```

## Why `strhelper?`

## Examples

`Code Smaples`

## Acknowledgement

Any other packages used

## Links

* [Go Slice Tricks Cheat Sheet](https://ueokande.github.io/go-slice-tricks/)
* [SliceTricks · golang/go Wiki](https://github.com/golang/go/wiki/SliceTricks)
* [Go Cheat Sheet](https://gist.github.com/ikennaokpala/b839c00e0123374e0b58)
* [The Go Programming Language Specification - The Go Programming Language](https://golang.org/ref/spec#Assignments)
* [go reflect - How to find the type of an object in Go?](https://stackoverflow.com/questions/20170275/how-to-find-the-type-of-an-object-in-go)
* [go - Can I compare variable types with .(type) in Golang?](https://stackoverflow.com/questions/34820469/can-i-compare-variable-types-with-type-in-golang)
* [reflection - Check type of struct in Go](https://stackoverflow.com/questions/45067382/check-type-of-struct-in-go)
* [Go - How can I check for type equality?](https://stackoverflow.com/questions/29444817/go-how-can-i-check-for-type-equality)
* [How to create dynamic items with reflection on interface type by astaxie/goorm](https://github.com/astaxie/goorm/blob/d35780c0d3cefa6f6171add5a864e7948e7fee8b/goorm.go#L207)
    * [struct to map function · astaxie/goorm](https://github.com/astaxie/goorm/blob/d35780c0d3cefa6f6171add5a864e7948e7fee8b/util.go#L111)

## Issues

- [Create your issues](https://gitlab.com/evatix-go/strhelper/-/issues)

## Notes

## Contributors

## License

[Evatix MIT License](/LICENSE)
