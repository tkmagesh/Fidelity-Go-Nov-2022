/*
	refactor into a function "generatePrimes" with (start and end) as the arguments returing a list prime numbers
	the main function prints the returned list from generatePrimes function
*/
package main

import "fmt"

func main() {
	primeNos := generatePrimes(3, 100)
	fmt.Println(primeNos)
}

func generatePrimes(start, end int) []int {
	// result := []int{}
	var result []int
LOOP:
	for no := start; no <= end; no++ {
		for i := 2; i < no; i++ {
			if no%i == 0 { // no is not a prime number
				continue LOOP
			}
		}
		result = append(result, no)
	}
	return result
}
