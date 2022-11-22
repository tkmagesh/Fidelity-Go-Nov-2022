package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == DivideByZeroError {
		fmt.Println("Please do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong.", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

/*
func divide(x, y int) (quotient int, remainder int, err error) {
	if y == 0 {
		err = errors.New("divisor cannot be zero")
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
*/

func divide(x, y int) (quotient int, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
