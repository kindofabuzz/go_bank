package comms

import (
	"fmt"
)

const name = "Jason"

func PresentOptions() {
	fmt.Printf("What would you like to do, %v?\n", name)
	fmt.Println()
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}
