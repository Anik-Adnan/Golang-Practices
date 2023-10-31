package main

import "fmt"

func main() {
	fmt.Println("enter your firstName")
	var name string
	fmt.Scan(&name) // one word

	fmt.Println("enter your age")
	var age int
	fmt.Scan(&age)
	
	fmt.Printf("FirstName : %v Age: %v", name, age)
	fmt.Println("enter your full name:")
	var fullname string
	fmt.Scanln()
	fmt.Scanln(&fullname) // one line
	fmt.Printf("Your full name: %v", fullname)
}
