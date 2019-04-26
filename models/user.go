package models

import (
	"database/sql"
	"log"
	"strings"
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

// GetUser gets user from the database
func GetUser(db *sql.DB, email string) User {
	var sb strings.Builder
	sb.WriteString("SELECT * FROM USERS ")
	sb.WriteString("WHERE email = $1 ")
	query := sb.String()

	row := db.QueryRow(query, email)
	user := User{}
	err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password)
	if err != nil {
		log.Fatal(err)
	}
	return user
}

// CreateUser functon creates a user
func CreateUser(db *sql.DB, email, username, password string) sql.Result {
	var sb strings.Builder
	sb.WriteString("INSERT INTO USERS ")
	sb.WriteString("(email, username, password) ")
	sb.WriteString("VALUES ($1, $2, $3) ")
	query := sb.String()

	res, err := db.Exec(query, email, username, password)
	if err != nil {
		log.Fatal(err)
	}
	return res

}
