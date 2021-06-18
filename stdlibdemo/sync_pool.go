package stdlibdemo

import "sync"

type Something struct {
	Name string
}

var pool = sync.Pool{
	New: func() interface{} {
		return &Something{}
	},
}

func main() {
	if s, isValid := pool.Get().(*Something); isValid {
		defer pool.Put(s)
		s.Name = "hello"
		// use the object
	}
}
