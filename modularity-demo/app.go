package main

import (
	"fmt"
	"modularity-demo/calculator"
)

func main() {
	fmt.Println("modularity-demo executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
}
