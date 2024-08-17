// PROMPT from chat
// Exercise: Building a Simple ATM Simulation
// Problem Description:
// Create a simple ATM simulation program that allows users to interact with a virtual bank account.
// The program should provide basic functionality such as checking the account balance, depositing money, withdrawing money, and exiting the program.
// The goal is to use control structures, functions, and variables effectively to manage user input and perform the necessary operations.

// Requirements:

// Initial Setup:
// Start with an initial balance of $1000.
// Display a menu with options for the user: Check Balance, Deposit Money, Withdraw Money, and Exit.

// Control Structures:
// Use a loop to repeatedly show the menu until the user decides to exit.
// Handle the user's menu selection using control structures like if-else or switch-case.
// Functions:

// Implement a function to check and display the current balance.
// Implement a function to deposit money into the account, updating the balance accordingly.
// Implement a function to withdraw money, ensuring the withdrawal amount does not exceed the current balance.
// Implement a function to display the menu and get the user's choice.
// Variables:

// Maintain a variable to keep track of the account balance.
// Use variables to handle user input and other necessary data for transactions.
// User Interaction:
// The user should be able to repeatedly interact with the ATM until they choose to exit.
// The user can check their balance, deposit funds, or withdraw funds based on their choice.
// The program should provide feedback on each transaction, such as confirming the new balance after a deposit or withdrawal.
// Goals:

// Reinforce the use of variables for tracking and updating the balance.
// Practice writing functions to organize the different parts of the program.
// Utilize control structures to manage the program’s flow and user interactions.
// This exercise will help you bring together concepts of control structures, functions, and variables into a cohesive program.

package main

import (
	"fmt"
)

func display() {
	fmt.Println("********** OPTIONS ****************\"n" +
		"SHOW BALANCE (type: show_balance)\n" +
		"DEPOSIT (type: deposit)\n" +
		"WITHDRAW (type: withdraw)\n" +
		"DISPLAY MENU (type: display)\n" +
		"QUIT (type: quit)")
}

func show_balance(balance *float64) {
	fmt.Printf("Current Balance: %f\n", *balance)
}

func deposite(balance *float64) {
	fmt.Printf("Balance: %f\n", *balance)
	fmt.Print("How much to deposite: ")
	var deposite float64
	fmt.Scanln(&deposite)
	*balance += deposite
	fmt.Printf("Transaction finished\n"+
		"Current Balance: %f\n", *balance)

}

func withdraw(balance *float64) {
	fmt.Printf("balance: %f\n", *balance)
	fmt.Print("How much to withdraw: ")
	var withdraw float64
	fmt.Scanln(&withdraw)

	if withdraw > *balance {
		fmt.Println("ERROR: withdraw amount is greater than balance")
		return
	}
	*balance -= withdraw
	fmt.Printf("Transaction finished\n"+
		"Current Balance: %f\n", *balance)
}

func main() {
	fmt.Println("Starting ATM Simulation")
	display()
	var current_amount float64 = 1000
	var option string
	for {
		fmt.Println("Enter option: ")
		fmt.Scanln(&option)

		switch option {
		case "show_balance":
			show_balance(&current_amount)
		case "deposit":
			deposite(&current_amount)
		case "withdraw":
			withdraw(&current_amount)
		case "display":
			display()
		case "quit":
			return
		}
	}

}
