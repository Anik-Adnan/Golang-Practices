package main

import (
	"fmt"
)

var pl = fmt.Println

// pass by reference
func changeValue(v *int) {
	*v += 5
}
func main() {
	p := 10
	var vPtr *int = &p
	pl("address of p   : ", &p) // address are same
	pl("address of vPtr: ", vPtr)
	pl("before value: ", p)
	changeValue(&p)
	pl("Change value :", p)

}
