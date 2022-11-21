package main

import "fmt"

// var x int = 100
var x = 100 // CAN have un-used variables at package scope

// x := 100

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/*
		//use type inference
		var msg = "Hello World!"
	*/

	msg := "Hello World!" // := syntax can be use ONLY in function (NOT at package scope)
	fmt.Println(msg)

	//var n int = 100 // CANNOT have un-used variables at function scope
}
