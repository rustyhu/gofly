package gointerface

import "log"

//Container is a generic container, accepting anything.
type Container []interface{}

//Put adds an element to the container.
func (c *Container) Put(elem interface{}) {
	*c = append(*c, elem)
}

//Get gets an element from the container.
func (c *Container) Get() interface{} {
	if len(*c) > 0 {
		elem := (*c)[0]
		*c = (*c)[1:]
		return elem
	}
	return nil
}

func TAExam() {
	var iCon = Container{}
	iCon.Put(7)
	iCon.Put(13)

	a, e := iCon.Get().(int)
	if !e {
		log.Println("Get failed!")
	} else {
		log.Println("Get a num:", a)
	}

}
