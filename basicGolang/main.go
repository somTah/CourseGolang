package main

import "fmt"

var numberInt, numberInt2 int = 100, 200
var msg string = "hello"

func main() {
	numberflaot := 25.6
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberflaot)
	fmt.Println(msg)

	fmt.Println(numberInt + numberInt2)
	fmt.Println(float64(numberInt) + numberflaot)
	fmt.Println(msg + "World")
	fmt.Println("my money =", numberInt)
}
