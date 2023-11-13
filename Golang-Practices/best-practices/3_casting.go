package main

import (
	"fmt"
	"reflect"
	"strconv"
)

var pl = fmt.Println

func main() {
	v1 := 1.5
	v2 := int(v1)
	pl(reflect.TypeOf(v2), v2)

	// string to int
	s1 := "121"
	s2, err := strconv.Atoi(s1)
	pl(s2, err, reflect.TypeOf(s2))

	// int to string
	i := 1212
	s := strconv.Itoa(i)
	pl(s, reflect.TypeOf(s))

	// string to float
	floatstring := "1223.245"
	if strToFloat, err := strconv.ParseFloat(floatstring, 64); err != nil {
		pl(strToFloat)
	}
}
