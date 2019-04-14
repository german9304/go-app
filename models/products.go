package models

// Product structure represents model product
type Product struct {
	Name        string
	Description string
}

// Elements variable
var Elements = []Product{
	Product{"Socks", "Nice socks, and are recommended!"},
	Product{"Jeans", "Nice Jeans, and are recommended!"},
}

// Data struct represents product
type Data struct {
	Products []Product
}

var DataModels = Data{Elements}
