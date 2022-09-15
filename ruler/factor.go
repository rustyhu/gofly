package ruler

///// Factors
type Factor interface {
	Calc() int64
	UpdateData(tr Transaction)
}

// Parameters are with Factor
type Parameter struct {
}

// TotalBalance is one type of detailed factor statistic data
type TotalBalance struct {
	Balance int
}
