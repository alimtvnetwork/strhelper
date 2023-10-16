package pk1

type (
	Requester interface {
		Field1() string
		SetField1(s string)
	}
	
	Outputter interface {
		Result() string
		SetResult(x string)
	}
)
