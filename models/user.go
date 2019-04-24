package models

import (
	"database/sql"
	"log"
)

//User struct represents user model
type User struct {
	ID       string
	Email    string
	Username string
	Password string
}

// GetAllUsers function gets all users from the database
func GetAllUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT * FROM USERS")
	if err != nil {
		log.Fatal(err)
	}
	var users []User
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	return users
}
