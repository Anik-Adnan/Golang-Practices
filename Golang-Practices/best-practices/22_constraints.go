package main

import (
	"fmt"
)

var pl = fmt.Println

type MyConstraint interface {
	int | float64
}

func getSumGeneric[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	pl("5 + 8 = ", getSumGeneric(5, 8))
	pl("5.5 + 8.3 = ", getSumGeneric(5.5, 8.3))
}
