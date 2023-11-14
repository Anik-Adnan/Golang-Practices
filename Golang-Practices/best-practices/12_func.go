package main

import (
	"fmt"
)

var pl = fmt.Println

func sayHello() {
	pl("Hello!!")
}
func getSumSUb(a, b int) (int, int) {
	return a + b, a - b
}
func getDiv(x float64, y float64) (res float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("you cann't divide by zero")
	} else {
		return x / y, nil
	}
}
func main() {
	a, b := 20, 10
	sum, sub := getSumSUb(a, b)
	sayHello()
	pl(getSumSUb(a, b))
	pl("Sum : ", sum)
	pl("Sub : ", sub)
	pl(getDiv(2, 0))

}
