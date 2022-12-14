package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	main - deferred")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	/*
		defer func() {
			fmt.Println("	f1 - deferred [1]")
		}()
		defer func() {
			fmt.Println("	f1 - deferred [2]")
		}()
		defer func() {
			fmt.Println("	f1 - deferred [3]")
		}()
	*/
	defer fmt.Println("	f1 - deferred [1]")
	defer fmt.Println("	f1 - deferred [2]")
	defer fmt.Println("	f1 - deferred [3]")
	fmt.Println("f1 started")
	f2Result := f2(100)
	fmt.Println("f2Result =", f2Result)
	fmt.Println("f1 completed")
}

func f2(x int) (result int) {
	defer func() {
		fmt.Println("	f2 - deferred")
		result = x * 5
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
	result = x * 2
	return
}
