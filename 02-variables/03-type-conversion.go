package main

import "fmt"

func main() {
	var i int32 = 100
	var f float64
	//use the typename like a function
	f = float64(i)
	fmt.Println(f)
}
