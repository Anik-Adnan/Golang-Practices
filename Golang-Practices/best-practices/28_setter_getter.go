package main

import (
	"fmt"
)

type User struct {
	name string
	age  int
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetAge(age int) {
	u.age = age
}

func (u *User) GetAge() int {
	return u.age
}

func main() {
	user := User{name: "Anik", age: 8}

	// Set the name using the setter method
	user.SetName("Adnan")

	// Get the name using the getter method
	name := user.GetName()
	fmt.Println("Name:", name)

	// Set the age using the setter method
	user.SetAge(40)

	// Get the age using the getter method
	age := user.GetAge()
	fmt.Println("Age:", age)

}
