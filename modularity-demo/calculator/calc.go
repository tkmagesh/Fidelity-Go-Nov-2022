package calculator

import "fmt"

var opCount int

func init() {
	fmt.Println("Calculator package initialized [1]")
}

func init() {
	fmt.Println("Calculator package initialized [2]")
}

func PrintStats() {
	fmt.Println("Operation Count :", opCount)
}

func Multiply(x, y int) int {
	return x * y
}
