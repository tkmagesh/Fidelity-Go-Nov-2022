package main

import (
	"fmt"
	"log"
	"time"
)

type OperationFn func(int, int)

func main() {

	logAdd := getLogOperation(add)
	profiledAdd := getProfiledOperation(logAdd)
	profiledAdd(100, 200)

	profiledSubtract := getProfiledOperation(getLogOperation(subtract))
	profiledSubtract(100, 200)

	getProfiledOperation(getLogOperation(func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	}))(100, 200)
}

func getProfiledOperation(fn OperationFn) OperationFn {
	return func(x, y int) {
		start := time.Now()
		fn(x, y)
		elapsed := time.Now().Sub(start)
		fmt.Println("Time Taken : ", elapsed)
	}
}

func getLogOperation(fn OperationFn) OperationFn {
	return func(x, y int) {
		log.Println("operation started")
		fn(x, y)
		log.Println("operation completed")
	}
}

//3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
