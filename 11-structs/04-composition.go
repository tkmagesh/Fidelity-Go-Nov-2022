package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Category string
}

type PerishableProduct struct {
	Product
	Expiry string
}

func main() {
	/*
		pen := Product{100, "Pen", 10, "Stationary"}
		fmt.Println(Format(pen))
		fmt.Println("After applying 10% discount")
		ApplyDiscount(&pen, 10)
		fmt.Println(Format(pen))
	*/

	// apple := PerishableProduct{Product{101, "Apple", 20, "Fruits"}, "5 Days"}
	apple := PerishableProduct{
		Product: Product{101, "Apple", 20, "Fruits"},
		Expiry:  "5 Days",
	}
	fmt.Println(apple)

	fmt.Println("Accessing the fields")
	fmt.Println(apple.Product.Id, apple.Product.Name, apple.Expiry)
	fmt.Println(apple.Name, apple.Expiry)

	//fmt.Println(Format(apple.Product))
	fmt.Println(FormatPerishableProduct(apple))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&apple.Product, 10)
	//fmt.Println(Format(apple.Product))
	fmt.Println(FormatPerishableProduct(apple))
}

func Format(product Product) string {
	return fmt.Sprintf("Id = %d, Name = %q, Cost = %f, Category = %q", product.Id, product.Name, product.Cost, product.Category)
}

func ApplyDiscount(product *Product, discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}

/* to include the Expiry in Format */
func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%s, Expiry = %q", Format(pp.Product), pp.Expiry)
}
