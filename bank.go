package main

import (
	"beast/bank/comms"
	"beast/bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println("--------------------")
	fmt.Println("Welcome to Go-Bank!!")
	fmt.Println("Give us a call yall!: ", randomdata.PhoneNumber())
	fmt.Println("--------------------")
	fmt.Println()

	for {

		comms.PresentOptions()

		var choice int
		fmt.Println()
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your Balance:  %.2f\n", accountBalance)

		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Must be greater than 0.")
				fmt.Println("----------------")
				continue
			}

			accountBalance += depositAmount

			fmt.Printf("Balance updated. New balance: %.2f\n", accountBalance)
			fmt.Println()
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Print("Withdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Must be greater than 0.")
				fmt.Println("----------------")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Nigga you ain't got that much.")
				fmt.Println("----------------")
				continue
			}

			accountBalance -= withdrawAmount

			fmt.Printf("Balance updated. New balance: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("See ya!")
			fmt.Println("Thank you for your patronage!")
			return

		}

	}

}
