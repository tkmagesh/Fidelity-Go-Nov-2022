package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) Increment() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {
	//runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go concurrentFn(wg)
	}
	wg.Wait()
	fmt.Println("counter :", counter.count)
}

func concurrentFn(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.Increment()
}
