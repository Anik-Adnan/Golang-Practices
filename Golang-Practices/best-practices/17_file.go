package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	f, err := os.Create("data.txt")
	checkNilError(err)
	defer f.Close()

	iPrimeArr := []int{2, 3, 5, 7}
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		checkNilError(err)

	}
	f, err = os.Open("data.txt")
	checkNilError(err)
	defer f.Close()
	scanFile := bufio.NewScanner(f)
	for scanFile.Scan() {
		pl("Prime: ", scanFile.Text())
	}
	checkNilError(scanFile.Err())
	// if err := scanFile.Err();err !=nil{
	// 	log.Fatal(err)
	// }

}
func checkNilError(err error) {
	if err != nil {
		// panic(err)
		log.Fatal(err)
	}
}
