package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() //scheduling the execution of the function through the Go scheduler (goroutine)
	f2()
	//block the execution of this function so that the go scheduler can execute the already scheduled goroutines
	fmt.Scanln() // DONOT do this in production
}

func f1() {
	fmt.Println("f1 invocation started")
	time.Sleep(3 * time.Second) // simulating a time consuming operation
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
