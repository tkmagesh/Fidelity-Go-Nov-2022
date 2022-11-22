/*
	refactor into a function "generatePrimes" with (start and end) as the arguments returing a list prime numbers
	the main function prints the returned list from generatePrimes function
*/
package main

import "fmt"

func main() {
LOOP:
	for no := 3; no <= 100; no++ {
		for i := 2; i < no; i++ {
			if no%i == 0 { // no is not a prime number
				continue LOOP
			}
		}
		fmt.Println("Prime No :", no)
	}
}
