package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func query(db *sql.DB) {

	var (
		id         int
		courseName string
		price      float64
		instructor string
	)
	var inputId int
	fmt.Scan(&inputId)
	query := "SELECT id, courseName, price, instructor FROM onlinecourse WHERE id = ?"
	if err := db.QueryRow(query, inputId).Scan(&id, &courseName, &price, &instructor); err != nil {
		log.Fatal(err)
	}
	fmt.Println(id, courseName, price, instructor)
}

func main() {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/database_name")
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golangcourse")
	if err != nil {
		fmt.Println("Failed to connect")
	} else {
		fmt.Println("Connect successfully")
	}
	query(db)
}
