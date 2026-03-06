package main

import (
	"beast/bank/comms"
	"beast/bank/fileops"
	"fmt"

	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"


func main() {
	dashes := "-----------------------------------------"

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
	}

	fmt.Println(dashes)
	fmt.Println("Welcome to Go-Bank!!")
	fmt.Printf("Give %v a call at %v\n", randomdata.SillyName(), randomdata.PhoneNumber())
	fmt.Print(dashes)
	fmt.Println()

	for {

		comms.PresentOptions()

		var choice int
		fmt.Println()
		fmt.Print("Enter choice: ")
		// fmt.Println("")
		fmt.Scanln(&choice)

		switch choice {
		case 0:
			continue
		case 1:
			fmt.Printf("\nYour Balance:  %.2f\n", accountBalance)

		case 2:
			fmt.Print("\nYour Deposit: \n")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Must be greater than 0.")
				fmt.Print(dashes, "\n")
				continue
			}

			accountBalance += depositAmount

			fmt.Printf("\nBalance updated. New balance: %.2f\n", accountBalance)
			fmt.Println()
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		case 3:
			fmt.Print("\nWithdraw amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("\nMust be greater than 0.")
				fmt.Print(dashes, "\n")
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("\nNigga you ain't got that much.")
				fmt.Print(dashes, "\n")
				continue
			}

			accountBalance -= withdrawAmount

			fmt.Printf("\nBalance updated. New balance: %.2f\n", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)

		default:
			fmt.Println("\nSee ya!")
			fmt.Println("Thank you for your patronage!")
			return

		}

	}

}
