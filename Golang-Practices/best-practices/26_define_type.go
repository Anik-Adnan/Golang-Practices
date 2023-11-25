package main

import (
	"fmt"
)

var pl = fmt.Println

type fl float64

func area(b fl, h fl) fl {
	return b * h
}

func main() {
	var base, height fl = 4.5, 8.6
	pl(area(base, height))
}
