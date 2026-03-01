package main

import (
	"fmt"
)

const name = "Jason"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to Go-Bank!!\n")

	for {

		fmt.Printf("What would you like to do, %v?\n", name)
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your Balance:  ", accountBalance)
		} else if choice == 2 {
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Must be greater than 0.")
				fmt.Println("----------------")
				continue
			}

			accountBalance += depositAmount
			fmt.Println("Balance updated. New balance: ", accountBalance)

		} else if choice == 3 {
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Must be greater than 0.")
				return
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Nigga you ain't got that much.")
				return
			}
			accountBalance -= withdrawAmount
			fmt.Printf("Balance updated. New balance: %.2f", accountBalance)

		} else {
			fmt.Println("See ya!")
			break

		}
	}

	fmt.Println("Thank you for your patronage!")
}
