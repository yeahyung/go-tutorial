package main

import (
	"fmt"
	"go-tutorial/tutorial/banking/banking"
	"log"
)

func main() {
	account := banking.NewAccount("yea")
	account.Deposit(100)
	fmt.Println(account.Balance())

	account.ChangeOwner("hyung")
	fmt.Println(account)

	err := account.Withdraw(200)
	if err != nil {
		// Fatalln kills program after print
		log.Fatalln(err)
	}
	fmt.Println(account.Balance())
}
