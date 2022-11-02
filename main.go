package main

import (
	"fmt"
	"os"
	"os/exec"
)

var balance float32 = 10000
var anotherTransaction int

func main() {
	fmt.Printf("\nWelcome to GTBank ATM\n")
	transaction()
}

func exit() {
	fmt.Println("\nThank you for banking with us!\nHave a nice day")
}
func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	err := c.Run()
	if err != nil {
		return
	}
}
func anotherTransact() {
	fmt.Printf("\n\nDo you want to perform another transaction?\nPress 1 to proceed and 2 to exit\n\n")
	_, err := fmt.Scan(&anotherTransaction)
	if err != nil {
		return
	}
	switch anotherTransaction {
	case 1:
		clear()
		transaction()
	default:
		exit()
	}
}
func transaction() {
	var userChoice int

	fmt.Printf("\nPlease select an option\n\n")
	fmt.Printf("1. Withdraw\n")
	fmt.Printf("2. Deposit\n")
	fmt.Printf("3. Balance\n")
	fmt.Printf("4. Exit\n\n")

	_, err := fmt.Scan(&userChoice)
	if err != nil {
		return
	}
	var amountToWithdraw float32
	var amountToDeposit float32

	switch userChoice {
	case 1: //Withdraw
		fmt.Printf("Please enter amount you want to withdraw: ")
		_, err := fmt.Scan(&amountToWithdraw)
		if err != nil {
			return
		}

		if amountToWithdraw > balance {
			fmt.Printf("\nInsufficient Funds")
			anotherTransact()
		} else {
			balance -= amountToWithdraw
			fmt.Printf("Withdrawal Successful\nYou have withdrawn %.2f NGN and your available balance is %.2f NGN",
				amountToWithdraw, balance)
			anotherTransact()
		}

	case 2: //Deposit
		fmt.Printf("Please enter amount to deposit: ")
		_, err := fmt.Scan(&amountToDeposit)
		if err != nil {
			return
		}

		balance += amountToDeposit

		fmt.Printf("Deposit successful!\nYour available balance is: %.2f NGN", balance)

		anotherTransact()

	case 3: //Check Balance
		fmt.Printf("\nYour available balance is: %.2f NGN", balance)
		anotherTransact()

	case 4: //Exit
		exit()
	}
}
