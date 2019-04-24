package models

import (
	"database/sql"
	"log"
	"strings"
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
	defer rows.Close()
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

// GetProduct gets pne product from database
func GetProduct(db *sql.DB, productID int) Product {
	row := db.QueryRow("SELECT id, name, likes, description, quantity FROM PRODUCT WHERE id = $1", productID)
	product := Product{}
	err := row.Scan(&product.ID, &product.Name, &product.Likes, &product.Description, &product.Quantity)
	if err != nil {
		log.Fatal(err)
	}
	return product
}

// InsertProduct inserts a row in product model
func InsertProduct(db *sql.DB, quantity, userID int, name, description string) sql.Result {
	var sb strings.Builder
	sb.WriteString("INSERT INTO PRODUCT (name, description, quantity, user_id) ")
	sb.WriteString("VALUES ($1, $2, $3, $4) ")
	query := sb.String()
	row, err := db.Exec(query, name, description, quantity, userID)
	if err != nil {
		log.Fatal(err)
	}
	return row
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
