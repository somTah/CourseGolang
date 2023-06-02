package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Id           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	e := employee{}
	err := json.Unmarshal([]byte(`{"Id": 101, "EmployeeName": "Sirasit", "Tel": "0258963258", "Email": "sirasit@gmail.com"}`), &e)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e.EmployeeName)
}
