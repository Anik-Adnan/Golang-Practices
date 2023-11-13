package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(10) + 1
	maxTries := 5

	for i := 0; i < maxTries; i++ {
		pl("Guess a number between 1 to 10 ")
		pl("Randm number is : ", randNum)
		reader := bufio.NewReader(os.Stdin)
		guessnum, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guessnum = strings.TrimSpace(guessnum)
		intGuess, err := strconv.Atoi(guessnum)
		if err != nil {
			log.Fatal(err)
		}
		if maxTries == i+1 {
			pl("too many tries!!")
			break
		}
		if intGuess > randNum {
			pl("Lower Value")
		} else if intGuess < randNum {
			pl("Upper value")
		} else {
			pl("Number Matched")
			break
		}

	}

}
