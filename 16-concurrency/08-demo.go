package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

func main() {
	//runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go concurrentFn(wg)
		// fn()
	}
	wg.Wait()
	fmt.Println("counter :", counter)
}

func fn() {
	counter++
}

func concurrentFn(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	{
		counter++
	}
	mutex.Unlock()
}
