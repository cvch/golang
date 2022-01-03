package example

type Example struct {
	PublicField  int
	privateField int
}

func NewStruct(a, b int) Example {
	return Example{
		PublicField:  a,
		privateField: b,
	}
}
