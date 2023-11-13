package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var a = [...]int{} // length is infered
	pl("infered array:", a)
	var arr [5]int // length is define
	arr[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5}

	fmt.Printf("%T %T\n", arr, arr2)
	pl(arr)

	for i := 0; i < len(arr2); i++ {
		fmt.Print(arr2[i], " ")
	}
	pl()

	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}
	arr3 := [2][]int{
		{1, 2},
		{3, 4},
	}
	pl(arr3, "\nlen :", len(arr3))
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

}
