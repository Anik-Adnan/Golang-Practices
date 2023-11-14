package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func main() {
	_, err := os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("file doesn't Exits")
	} else {
		f, err := os.OpenFile("data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 011)
		checkNilError(err)
		defer f.Close()
		_, err = f.WriteString("11\n")
		checkNilError(err)

	}

}
func checkNilError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
