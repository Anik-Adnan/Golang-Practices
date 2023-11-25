package main

import (
	"fmt"
)

type customer struct {
	name    string
	address string
	balance float64
}

func getCustomerInfo(c customer) {
	pl("Customer name: ", c.name, " Address : ", c.address)
}
func newCustomerAddress(c *customer, address string) {
	c.address = address
}

var pl = fmt.Println

func main() {
	var c1 customer
	c1.name = "Anik Adnan"
	c1.address = "Dhaka"
	c1.balance = 50000.50
	pl(c1)
	getCustomerInfo(c1)
	newCustomerAddress(&c1, "Tangail")
	getCustomerInfo(c1)

	c2 := customer{"Munna Pervez", "Meherpur", 25.80}
	pl(c2)

}
