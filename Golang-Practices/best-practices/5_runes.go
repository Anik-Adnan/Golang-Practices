package main

import (
	"fmt"
	"unicode/utf8"
)

var pl = fmt.Printf

func main() {
	runeStr := "abcABC"
	pl("Rune Count: ", utf8.RuneCountInString(runeStr))

	for i, runeVal := range runeStr {
		pl("%d : %#U : %c\n", i, runeVal, runeVal)
	}

}
