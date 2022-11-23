package main

import "fmt"

func main() {
	// var x interface{} //go 1.18
	var x any
	x = 100
	x = "This is a string"
	x = true
	// x = 100.99
	// x = struct{}{}
	// x = []int{3, 1, 4, 2, 5}
	fmt.Println(x)

	// x = 100
	// x = "This is a string"
	// y := x + 200
	// y := x.(int) + 200 //not runtime safe
	if val, ok := x.(int); ok { // runtime safe
		y := val + 200
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}

	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 200 =", val+200)
	case float64:
		fmt.Println("x is a float64, 10% of x =", val*(0.1))
	case string:
		fmt.Println("x is a string, len(x) =", len(val))
	case struct{}:
		fmt.Println("x is an empty struct")
	case []int:
		fmt.Println("x is a slice of int")
	default:
		fmt.Println("unknown type")

	}

}
