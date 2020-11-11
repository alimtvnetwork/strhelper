package testscore

type TestCaseMessager interface {
	FuncName() string
	Value() interface{}
	Expected() interface{}
	Actual() interface{}
}
