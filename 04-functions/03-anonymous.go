/*
	anonymous functions
		 => functions with no names and immediately invoked
*/

package main

import "fmt"

func main() {
	func() {
		fmt.Println("hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	msg := func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a nice day!", userName)
	}("Suresh")
	fmt.Println(msg)

	addResult := func(x, y int) int {
		return x + y
	}(100, 200)
	fmt.Println("Add Result =", addResult)

	q, r := func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
}
