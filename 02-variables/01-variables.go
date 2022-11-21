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
	msg = "Hello everyone!"
	fmt.Println(msg)

	//var n int = 100 // CANNOT have un-used variables at function scope

	//declaring multiple variables and initializing them
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of 100 and 200 is"
		var result int = x + y
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of 100 and 200 is"
		var result int = x + y
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of 100 and 200 is"
			result int    = x + y
		)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Sum of 100 and 200 is"
			result = x + y
		)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of 100 and 200 is"
			result    = x + y
		)
	*/

	x, y, str := 100, 200, "Sum of 100 and 200 is"
	result := x + y
	fmt.Println(str, result)
}
