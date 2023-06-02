package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func creatingTable(db *sql.DB) {

	query := `CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`

	if _, err := db.Exec(query); err != nil {
		log.Fatal(err)
	}
}

func insert(db *sql.DB) {

	var username string
	var password string
	fmt.Scan(&username)
	fmt.Scan(&password)
	createAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createAt)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	fmt.Println(id)
}
func deleteUser(db *sql.DB) {

	var deleteId int
	fmt.Scan(&deleteId)
	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, deleteId)
	if err != nil {
		log.Fatal(err)
	}
}

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
	// creatingTable(db)
	// insert(db)
	deleteUser(db)
}
