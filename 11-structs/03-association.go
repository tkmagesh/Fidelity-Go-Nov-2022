package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
	Org    *Organization
}

type Organization struct {
	Name string
	City string
}

func main() {
	/*
		e1 := Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 1000,
			Org: Organization{
				Name: "Fidelity",
				City: "Bengaluru",
			},
		}
		fmt.Println(e1.Org.Name, e1.Org.City)
	*/
	fidelity := &Organization{
		Name: "Fidelity",
		City: "Bengaluru",
	}
	e1 := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 1000,
		Org:    fidelity,
	}
	// fmt.Println(e1.Org.Name, e1.Org.City)
	fmt.Println(e1)
	e2 := Employee{
		Id:     200,
		Name:   "Suresh",
		Salary: 20000,
		Org:    fidelity,
	}
	fmt.Println(e2)

	fmt.Println("After changing the Org city....")
	fidelity.City = "Gurugram"
	fmt.Println("Org =", fidelity)
	fmt.Println("e1.Org =", e1.Org)
	fmt.Println("e2.Org =", e2.Org)
}
