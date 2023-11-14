package main

import (
	"fmt"
)

var pl = fmt.Println

// varadic function
func getTotal(nums ...int) int64 {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return int64(sum)
}

// pass by value
func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	pl(getTotal(1, 2, 3, 4, 5, 6))
	vArr := []int{1, 2, 3, 4, 5, 6}
	pl(getArraySum(vArr))

}
