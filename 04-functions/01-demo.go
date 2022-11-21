/* function basics */
package main

import "fmt"

func main() {
	say_hi()
}

//01 - basics
func say_hi() {
	fmt.Println("Hi there!")
	greet("Magesh")
	fmt.Println(getGreetMsg("Suresh"))
	fmt.Println("100 + 200 =", add(100, 200))

	// fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	//printing only the quotient
	/*
		q, r := divide(100, 7) // error : r declared by not used
		fmt.Printf("Dividing 100 by 7, quotient = %d \n", q)
	*/

	/*
		q := divide(100, 7) // error : assignment mismatch: 1 variable but divide returns 2 values
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/

	q, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
}

//02 - with parameters
func greet(userName string) {
	fmt.Printf("Hi %s, Have a nice day!\n", userName)
}

//03 - with return value
func getGreetMsg(userName string) string {
	return fmt.Sprintf("Hi %s, Have a nice day!", userName)
}

//04 - with more than 1 parameter
/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

//05- with more than 1 return results
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

//06 - using named return results
func divide(x, y int) (quotient, remainder int) { /* variables declared and initilized with the default value of 'int' ie., 0 */
	quotient = x / y
	remainder = x % y
	return
}
