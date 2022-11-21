/*
	Write a program that prints all the prime numbers between 3 and 100
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
