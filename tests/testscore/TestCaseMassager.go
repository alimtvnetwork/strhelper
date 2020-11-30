package testscore

type TestCaseMessenger interface {
	FuncName() string
	Value() interface{}
	Expected() interface{}
	Actual() interface{}
}
