package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	arrStr := "abcde"
	runeArr := []rune(arrStr)

	for _, v := range runeArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	byteArr := []byte{'a', 'n', 'i', 'k'}
	bstr := string(byteArr[:])
	pl("I'm a byte-string:", bstr)
}
