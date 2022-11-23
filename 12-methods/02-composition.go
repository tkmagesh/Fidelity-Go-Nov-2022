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

type PerishableProduct struct {
	Product
	Expiry string
}

//overriding
func (pp PerishableProduct) Format() string {
	return fmt.Sprintf("%s, Expiry = %q", pp.Product.Format(), pp.Expiry)
}

func main() {

	apple := PerishableProduct{
		Product: Product{101, "Apple", 20, "Fruits"},
		Expiry:  "5 Days",
	}
	fmt.Println(apple)

	fmt.Println("Accessing the fields")
	fmt.Println(apple.Product.Id, apple.Product.Name, apple.Expiry)
	fmt.Println(apple.Name, apple.Expiry)

	//fmt.Println(Format(apple.Product))
	fmt.Println(apple.Format())
	fmt.Println("After applying 10% discount")
	apple.ApplyDiscount(10)
	fmt.Println(apple.Format())
}

/* to include the Expiry in Format */
