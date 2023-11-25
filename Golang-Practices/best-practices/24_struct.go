package main

import (
	"fmt"
)

var pl = fmt.Println

type rectangle struct {
	length, heigth float64
}

func (r rectangle) Area() float64 {
	return r.heigth * r.length
}
func main() {
	r1 := rectangle{5, 10}
	pl("Rectangle area : ", r1.Area())
}
