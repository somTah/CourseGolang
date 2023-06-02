package main

import "fmt"

var score int

func main() {
	// myMoney := 101
	// if myMoney > 100 {
	// 	fmt.Println("yes")
	// } else {
	// 	fmt.Println("no")
	// }
	fmt.Println("grade calculator")
	fmt.Scanf("%d", &score)
	fmt.Println("your grade", score)
	grade(score)
}
func grade(score int) {
	myGrade := score
	if myGrade >= 80 {
		fmt.Println("A")
	} else if myGrade >= 70 {
		fmt.Println("B")
	} else if myGrade >= 60 {
		fmt.Println("C")
	} else if myGrade >= 50 {
		fmt.Println("D")
	} else {
		fmt.Println("F")
	}
}
