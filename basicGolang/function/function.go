package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func plus(val1 int, val2 int) int {
	// result := val1 + val2
	// fmt.Println("result =", result)
	return val1 + val2
}

func plusMuli(val1, val2, val3 int) int {
	return val1 + val2 + val3
}
func main() {
	hello()
	result := plus(5, 2)
	fmt.Println("result =", result)
	result2 := plusMuli(2, 5, 6)
	fmt.Println("rs =", result2)
}
