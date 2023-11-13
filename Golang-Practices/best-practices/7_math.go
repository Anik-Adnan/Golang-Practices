package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	//random value
	seedSecs := time.Now().Unix()
	pl(seedSecs)
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1 //0 to 50 random number
	pl("Random number :", randNum)

	//some math operations
	pl("Abs(-123) : ", math.Abs(-123))
	pl("Pow(5,2)  : ", math.Pow(5, 2))
	pl("Sqrt(25)  : ", math.Sqrt(25))
	pl("Cbrt(8)   : ", math.Cbrt(8))
	pl("Ceil(8.4) : ", math.Ceil(8.4))
	pl("Floor(8.6) : ", math.Floor(8.6))
	pl("Round(8.6) : ", math.Round(8.6))
	pl("log2(8)    : ", math.Log2(8))
	pl("Log10(100) : ", math.Log10(100))
	pl("Log(2) 	  : ", math.Log(2))
	pl("Max(5,10) : ", math.Max(5, 10))
	pl("Min(5,10) : ", math.Min(5, 10))

	// degree to radian
	r90 := 90 * math.Pi / 180
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%.4f radiands = %.2f degrees ", r90, d90)

	pl("\nSin(90)   = ", math.Sin(r90))

}
