/*
Create a product struct with the following fields
	Id
	Name
	Cost
	Category

Create the following functions
	Format() => returns a formatted string of the given product as below:
		" Id = 100, Name = "pen", Cost = 10.50, Category = "stationary" "
	ApplyDiscount() => update the given product's cost after applying the given discount(%)
*/

package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

func main() {
	pen := Product{100, "Pen", 10, "Stationary"}
	fmt.Println(Format(pen))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&pen, 10)
	fmt.Println(Format(pen))
}

func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %f, Category = %q", product.Id, product.Name, product.Cost, product.Category)
}

func ApplyDiscount(product *Product, discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
