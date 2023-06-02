package main

import "fmt"

var product [4]string
var price [4]float32

func main() {
	product[0] = "Macbook"
	product[1] = "iPad"
	product[2] = "iPhone"
	product[3] = "AirPods"
	price := [4]float32{40000, 30000, 20000, 2000}
	fmt.Println(product)
	fmt.Println(price)
}
