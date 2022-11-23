/*
Conver the "Format()" and "ApplyDiscount" into METHODS of Product
*/

package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %f, Category = %q", product.Id, product.Name, product.Cost, product.Category)
}

func (product *Product) ApplyDiscount(discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}

func main() {
	pen := Product{100, "Pen", 10, "Stationary"}
	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	// ApplyDiscount(&pen, 10)
	// (&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10) //pen is converted to a pointer (receiver expects a pointer) automatically
	// fmt.Println(Format(pen))
	fmt.Println(pen.Format())
}
