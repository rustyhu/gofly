package pattern

import (
	"testing"
	// "github.com/stretchr/testify/assert"
)

func TestBasicDerive(t *testing.T) {
	p1 := Base{&Product1{val: 2}}
	p1.Run()
}
