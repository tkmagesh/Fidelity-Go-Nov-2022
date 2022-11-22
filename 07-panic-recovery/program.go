package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError error = errors.New("divisor cannot be zero")

func main() {
	defer func() {
		fmt.Println("deferred - main")
		err := recover()
		if e, ok := err.(error); ok {
			if e == DivideByZeroError {
				fmt.Println("attempted to divide by zero")
				return
			}
			if e != nil {
				fmt.Println("exiting because of a panic", err)
				return
			}
		}
		fmt.Println("Thank you!")
	}()
	divisor := 0
	/*
		q, r := divide(100, divisor)
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	*/

	q, r, err := divideClient(100, divisor)
	if err == DivideByZeroError {
		fmt.Println("Do not attempt to divide by zero")
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		//converting a panic to an error
		if er := recover(); er != nil {
			err = er.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party library - we dont have any control over this code
func divide(x, y int) (int, int) {
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
