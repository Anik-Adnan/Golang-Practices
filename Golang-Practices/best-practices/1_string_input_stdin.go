package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var pl = fmt.Print

func main() {
	pl("What is your name?\n")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		pl("Hello! ", name)
	}
	// formated string template
	fs := fmt.Sprintf("Hello ! %s\n", name)
	pl(fs)

}
