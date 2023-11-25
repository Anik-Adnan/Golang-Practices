package main

import (
	"fmt"
)

type Account struct {
	balance int
}

func (a *Account) Deposit(amount int) {
	a.balance += amount
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.balance {
		return fmt.Errorf("Insufficient funds")
	}

	a.balance -= amount
	return nil
}

func main() {
	account := Account{balance: 100}

	account.Deposit(50)
	fmt.Println("Balance:", account.balance) // Output: Balance: 150

	err := account.Withdraw(75)
	if err != nil {
		fmt.Println(err) // Output: Insufficient funds
	} else {
		fmt.Println("Balance:", account.balance) // Output: Balance: 75
	}
}
