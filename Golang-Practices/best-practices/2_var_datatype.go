package main

import (
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	// var is camelcase
	var FName string = "Anik"
	var LName string = "Adnan"

	var v1, v2 = 10, 12
	v3 := "adnan"
	var v4 = 500.45
	pl(FName, LName, v1, v2, v3, v4)
	pl(v1, v2)
	pl(v3, v4) //concate

	pl(reflect.TypeOf(10))
	pl(reflect.TypeOf(10.245))
	pl(reflect.TypeOf("Anik"))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf('c')) //rune int32

}
