package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Println("%v", promt)
	input, _ := reader.ReadString('\n')
	value, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		message, _ := fmt.Scanf("%v must number only", promt)
		panic(message)
	}
	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}
func subtract(value1, value2 float64) float64 {
	return value1 - value2
}
func multiply(value1, value2 float64) float64 {
	return value1 * value2
}
func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string {
	fmt.Println("operator is (+ - * /)")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}
func main() {
	var loop int
	fmt.Println("คุณต้องการคำนวณกี่ตัวเลข: ")
	fmt.Scan(&loop)
	for i := 0; i < loop; i++ {
		fmt.Println("Number", i)
		fmt.Println("Operator :")
	}
	value1 := getInput("value1 = ")
	value2 := getInput("value2 = ")

	var result float64

	switch operator := getOperator(); operator {
	case "+":
		result = add(value1, value2)
	case "-":
		result = subtract(value1, value2)
	case "*":
		result = multiply(value1, value2)
	case "/":
		result = divide(value1, value2)
	default:
		panic("wrong operator")
	}
	fmt.Println("result is ", result)
}
