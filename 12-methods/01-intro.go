package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

/* receiver - the type for which the function should act as a method */
func (e Employee) WhoAmI() {
	fmt.Printf("I am an Employee (Name : %q)\n", e.Name)
}

func main() {
	emp := Employee{100, "Magesh", 10000}
	emp.WhoAmI()
}
