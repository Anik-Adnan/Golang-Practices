package main

import (
	person "example/module-practice/mypackage"
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	pl("hello ", person.Name)
	intArr := []int{2, 3, 5, 7, 11, 13}
	strArr := person.IntArrToStrArr(intArr)
	pl(strArr)
	pl(reflect.TypeOf(strArr))
}
