package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var person map[string]int64
	person = make(map[string]int64)
	person["Anik Adnan"] = 24
	person["Farhad"] = 25
	person["Md. Saif"] = 26
	pl("Anik Adnan's age: ", person["Anik Adnan"])
	pl("Munna's age: ", person["Munna"]) // 0
	pl("print name and age")
	for key, val := range person {
		pl("Name: ", key, " and Age : ", val)
	}
	fmt.Println("delete a name")
	delete(person, "Farhad")
	pl("farhad name is deleted", person)

}
