package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Id           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Sirasit", "0258963258", "sirasit@gmail.com"})
	fmt.Println(string(data))
}
