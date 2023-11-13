package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	sl1 := make([]string, 2)
	sl1[0] = "anik"
	sl1[1] = "adnan"
	pl("slice :", sl1)
	fmt.Printf("Type %T\n", sl1)
	pl("slice size: ", len(sl1), "capacity: ", cap(sl1))

	arr := []string{1: "hello", 2: "hi"}
	sl1 = append(sl1, arr...)
	pl(sl1)
	pl("slice size: ", len(sl1), "capacity: ", cap(sl1))

}
