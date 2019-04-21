package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	_ "github.com/lib/pq" // imported for use postgresql
	"github.com/shopcart/apiserver"
)

type userdb struct {
	User     string
	Password string
}

const (
	dbname = "mydb"
	port   = 5432
)

var (
	fileContent []byte
	user        userdb
	db          *sql.DB
	err         error
)

func readFile() []byte {
	cur, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fileJSON := path.Join(cur, "db", "userinfo.json")
	fileContent, err := ioutil.ReadFile(fileJSON)
	if err != nil {
		log.Fatal(err)
	}
	return fileContent
}

func init() {
	fileContent = readFile()
	err = json.Unmarshal(fileContent, &user)
	log.Println(user)
	connStr := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable port=%v",
		user.User, dbname, user.Password, port)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("init db")
}

// InitDbApp function is used to store db conntection
// to gloabl variable
func InitDbApp() {
	apiserver.Global["db"] = db
}
