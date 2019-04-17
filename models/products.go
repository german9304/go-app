package models

// Product structure represents model product
type Product struct {
	Name        string
	Description string
	Likes       int
	Quantity    int
}

// Elements variable
var Elements = []Product{
	Product{Name: "Socks", Description: "Nice socks, and are recommended!", Quantity: 10},
	Product{Name: "Jeans", Description: "Nice Jeans, and are recommended!", Quantity: 23},
}

// Data struct represents product
type Data struct {
	Products []Product
}

// DataModels reprensets model from the app
var DataModels = Data{Elements}
