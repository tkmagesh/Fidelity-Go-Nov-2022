package main

import (
	"fmt"

	"modularity-demo/calculator"
	"modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	color.Red("modularity-demo executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println(calculator.Multiply(100, 200))
	calculator.PrintStats()
	fmt.Println(utils.IsEven(21))
}
