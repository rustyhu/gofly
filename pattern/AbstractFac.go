package pattern

// ref to https://coolshell.cn/articles/8961.html#GoF%E7%9A%8423%E4%B8%AA%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F

type FacSeries uint8

const (
	FacA FacSeries = iota
	FacB
	FacC
	FacD
	FacE
)

type Prod interface {
}

// OOP implementation
type MyAbstractFactory interface {
	SetFactory()
	GetProd() Prod
}
type myAbstractFac map[FacSeries]MyAbstractFactory

// It is not necessary to implement design pattern always in OOP

/********** data driven simplified implementation **********/

type ProdSeries uint8

const (
	EProdA ProdSeries = iota
	EProdB
	EProdC
)

type SimpleFac func() Prod

var myFactoryProducer map[ProdSeries]SimpleFac

type ProdA struct {
}

func AFExam() {
	// register
	myFactoryProducer[EProdA] = func() Prod {
		return ProdA{}
	}

}
