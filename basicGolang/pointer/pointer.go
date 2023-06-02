package main

import "fmt"

func zeroVal(iVal int) {
	iVal = 0
}

func zeroPointer(iPointer *int) {
	*iPointer = 0
}
func main() {
	i := 1
	fmt.Println("i =", i)
	zeroVal(i)
	fmt.Println("i from function", i)
	zeroPointer(&i)
	fmt.Println("i from function pointer", i)
	fmt.Println("i  address from function pointer", &i)
}
