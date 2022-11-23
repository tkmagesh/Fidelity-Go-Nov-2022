package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// var wg *sync.WaitGroup = &sync.WaitGroup{}
	wg := &sync.WaitGroup{}

	wg.Add(1) // increment the wg counter by 1
	go f1(wg)
	f2()
	wg.Wait() // block until the wg counter becomes 0
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrement the wg counter by 1
	fmt.Println("f1 invocation started")
	time.Sleep(3 * time.Second) // simulating a time consuming operation
	fmt.Println("f1 invocation completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
