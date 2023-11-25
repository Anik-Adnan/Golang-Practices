package main

import (
	"fmt"
)

var pl = fmt.Println

type business struct {
	name    string
	address string
	contact
}
type contact struct {
	fName, lName, phone string
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s ", b.name, b.contact.fName, b.contact.lName)
}
func main() {
	contact1 := contact{"Anik", "Adnan", "0799-7579921"}
	busn1 := business{
		"Food Corner",
		"124,Mirpur,Dhaka",
		contact1,
	}
	busn1.info()

}
