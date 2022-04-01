package ctrlsys

///// Rule structure
type Rule interface {
	UpdateData()
	Check() bool
}

type Factor struct {
}

///// Rule structure END
