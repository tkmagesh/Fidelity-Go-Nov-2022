/*
	Apply error handling & refactor to use custom errors
*/

package main

import (
	"errors"
	"fmt"
)

var InvalidUserChoiceError = errors.New("invalid user choice")

func main() {
	var n1, n2 int
	for {
		userChoice, err := getUserChoice()
		if err != nil {
			fmt.Println(err)
			continue
		}
		if userChoice == 5 {
			fmt.Println("Thank you")
			break
		}
		operation, err := getOperationByUserChoice(userChoice)
		if err == InvalidUserChoiceError {
			fmt.Println("Invalid choice")
			continue
		}
		n1, n2 = getOperands()
		process(operation, n1, n2)
	}
}

func process(operation func(int, int) int, n1, n2 int) error {
	result := operation(n1, n2)
	fmt.Println("Result :", result)
	return nil
}

func getOperationByUserChoice(userChoice int) (operation func(int, int) int, err error) {
	switch userChoice {
	case 1:
		operation = add
	case 2:
		operation = subtract
	case 3:
		operation = multiply
	case 4:
		operation = divide
	default:
		err = InvalidUserChoiceError
	}
	return
}

func getUserChoice() (userChoice int, err error) {
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")
	fmt.Println("Enter your choice :")
	_, err = fmt.Scanln(&userChoice)
	return
}

func getOperands() (n1, n2 int) {
	// https://stackoverflow.com/questions/14640218/how-to-flush-stdin-after-fmt-scanf-in-go
	for {
		fmt.Println("Enter 2 numbers :")
		_, err := fmt.Scanln(&n1, &n2)
		if err != nil {
			fmt.Println("Error parsing the input. Please enter again")
			continue
		}
		return
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
