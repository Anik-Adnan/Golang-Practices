package main

import (
	"fmt"
)

var pl = fmt.Println

// pass by value
func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// pass by reference
func changeValue(v *int) {
	*v += 5
}
func main() {
	vArr := []int{1, 2, 3, 4, 5, 6}
	pl(getArraySum(vArr))

	val := 10
	pl("before value: ", val)
	changeValue(&val)
	pl("Change value :", val)
}
