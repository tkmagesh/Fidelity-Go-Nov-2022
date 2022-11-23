package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var noOfGoroutines int
	wg := &sync.WaitGroup{}

	flag.IntVar(&noOfGoroutines, "count", 0, "number of goroutines to execute")
	flag.Parse()

	rand.Seed(7)
	fmt.Printf("# of goroutines : %d\n", noOfGoroutines)

	fmt.Println("Hit ENTER to start....")
	fmt.Scanln()
	for i := 1; i <= noOfGoroutines; i++ {
		wg.Add(1)
		delay := rand.Intn(20)
		go fn(i, time.Duration(delay)*time.Second, wg)
	}
	wg.Wait()
	fmt.Println("Exiting main")
	fmt.Println("Hit ENTER to stop....")
	fmt.Scanln()
}

func fn(id int, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("fn[%d] started - wait duration = %v\n", id, delay)
	time.Sleep(delay)
	fmt.Printf("fn[%d] completed\n", id)
}
