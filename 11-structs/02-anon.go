package main

import "fmt"

func main() {
	// var s struct{}
	s := struct{}{}
	fmt.Println(s)

	emp := struct {
		Id     int
		Name   string
		Salary float32
	}{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	fmt.Println(emp)
	PrintEmployee(emp)
}

func PrintEmployee(e struct {
	Id     int
	Name   string
	Salary float32
}) {
	fmt.Printf("Id = %d, Name = %q, Salary = %f\n", e.Id, e.Name, e.Salary)
}
