package main

import "fmt"

type Employee struct {
	Id     int
	Name   string
	Salary float32
}

func main() {
	// var emp Employee
	// var emp Employee = Employee{100, "Magesh", 10000}
	// var emp Employee = Employee{Id: 100, Name: "Magesh", Salary: 10000}
	/*
		var emp Employee = Employee{
			Id:     100,
			Name:   "Magesh",
			Salary: 10000,
		}
	*/
	emp := Employee{
		Id:     100,
		Name:   "Magesh",
		Salary: 10000,
	}
	// fmt.Printf("%#v\n", emp)
	fmt.Printf("%+v\n", emp)

	fmt.Println("Accessing fields")
	fmt.Printf("Id = %d, Name = %q, Salary = %.2f\n", emp.Id, emp.Name, emp.Salary)

	//creating a copy
	/*
		emp2 := emp
		emp.Salary = 20000
		fmt.Println("emp =", emp)
		fmt.Println("emp2 =", emp2)
	*/

	var emp2 *Employee = &emp
	emp.Salary = 20000
	/*
		fmt.Println("emp =", emp)
		fmt.Println("emp2 =", emp2)
	*/
	fmt.Printf("emp.Salary = %.2f\n", emp.Salary)
	// fmt.Printf("emp2.Salary = %.2f\n", (*emp2).Salary)
	fmt.Printf("emp2.Salary = %.2f\n", emp2.Salary)

	emptyEmpPtr := new(Employee) // pointer to an default Employee object
	fmt.Println("emptyEmpPtr =", emptyEmpPtr)

}
