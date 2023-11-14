package main

import (
	"fmt"
)

var pl = fmt.Println

func passingArray(pArr *[4]int) {
	for i := 0; i < len(*pArr); i++ {
		pArr[i] += 5
	}

}

// slice always pass by reference
func getSlicesAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))
	for _, num := range nums {
		sum += num
	}
	nums[2] = 500
	return (sum / numSize)
}
func main() {
	pArr := [4]int{1, 2, 3, 4}
	passingArray(&pArr)
	pl(pArr)

	mySlice := []float64{5, 10, 15}
	pl("Average : ", getSlicesAverage(mySlice...))
	pl(mySlice[2])

}
