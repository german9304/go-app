package models

// Product structure represents model product
type Product struct {
	Name        string
	Description string
	Likes       int
}

// Elements variable
var Elements = []Product{
	Product{"Socks", "Nice socks, and are recommended!", 10},
	Product{"Jeans", "Nice Jeans, and are recommended!", 23},
}

// Data struct represents product
type Data struct {
	Products []Product
}

var DataModels = Data{Elements}
