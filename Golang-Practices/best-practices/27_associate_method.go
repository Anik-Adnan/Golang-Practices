package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func (e Employee) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", e.name, e.age)
}

func main() {
	e := Employee{name: "Anik", age: 30}
	e.Greet()
}
