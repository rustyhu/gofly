package gointerface

type Baif interface {
	sfunc(int) int
}

type Drif interface {
	Baif

	dfunc(int)
}

type mySt struct {
	id int
}

func (m *mySt) sfunc(int) int {
	// satisfy baif interface
	return 3 + m.id
}

func (m *mySt) dfunc(int) {
}

// Interface dependency
// var _ Baif = (*mySt)(nil)
var _ Drif = (*mySt)(nil)
