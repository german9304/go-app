package models

import (
	"database/sql"
	"log"
)

// Product structure represents model product
type Product struct {
	ID          int
	Name        string
	Description string
	Likes       int
	Quantity    int
}

// GetAllProducts gets all products from database
func GetAllProducts(db *sql.DB) []Product {
	rows, err := db.Query("SELECT id, name, likes, description, quantity FROM PRODUCT")
	if err != nil {
		log.Fatal(err)
	}
	var products []Product
	for rows.Next() {
		product := Product{}
		err = rows.Scan(&product.ID, &product.Name, &product.Likes, &product.Description, &product.Quantity)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	return products
}

// Elements variable
var Elements = []Product{
	Product{ID: 1, Name: "Socks", Description: "Nice socks, and are recommended!", Quantity: 10},
	Product{ID: 2, Name: "Jeans", Description: "Nice Jeans, and are recommended!", Quantity: 23},
}

// Data struct represents product
type Data struct {
	Products []Product
}

// DataModels reprensets model from the app
var DataModels = Data{Elements}
