package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println("product = ", product)

	//add
	product["Macbook"] = 40000
	product["iPhone"] = 30000
	product["iPad"] = 25000
	fmt.Println(product)
	//delete
	delete(product, "iPad")
	fmt.Println(product)

	//update
	product["iPhone"] = 20000
	fmt.Println(product)

	value := product["iPhone"]
	fmt.Println(value)

	courseName := map[string]string{"101": "Java", "102": "C#"}
	fmt.Println("", courseName)
}
