package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

const (
	username = "chiennm"
	password = "chien0110@123"
	hostname = "127.0.0.1:3306"
	dbname   = "ecommerce"
)

func main() {

	db, err := sql.Open("mysql", dsn(""))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
		return
	} else {
		fmt.Println("ok")
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	} else {
		fmt.Println("Successful database connection")
	}
}
func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
