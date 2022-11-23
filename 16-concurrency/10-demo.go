package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	//runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go concurrentFn(wg)
	}
	wg.Wait()
	fmt.Println("counter :", counter)
}

func concurrentFn(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1)
}
