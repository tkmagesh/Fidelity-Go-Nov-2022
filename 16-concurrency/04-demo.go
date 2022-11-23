package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(1) // increment the wg counter by 1
	go f1()

	wg.Add(1)
	go f2()
	wg.Wait() // block until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("	f1 invocation started")
	time.Sleep(3 * time.Second) // simulating a time consuming operation
	fmt.Println("	f1 invocation completed")
	wg.Done() // decrement the wg counter by 1
}

func f2() {
	fmt.Println("	f2 invocation started")
	time.Sleep(5 * time.Second) // simulating a time consuming operation
	fmt.Println("	f2 invocation completed")
	wg.Done() // decrement the wg counter by 1
}
