package main

import "fmt"

type employee struct { //tys
	employeeId   string
	employeeName string
	phone        string
}

func main() {
	// employee1 := employee{
	// 	employeeId:   "101",
	// 	employeeName: "Pradoo",
	// 	phone:        "0258963254",
	// }
	// fmt.Println("employee =", employee1)
	// employeeList := [3]employee{}
	// employeeList[0] = employee{
	// 	employeeId:   "101",
	// 	employeeName: "Pradoo",
	// 	phone:        "0258963258",
	// }
	// employeeList[1] = employee{
	// 	employeeId:   "102",
	// 	employeeName: "Prayad",
	// 	phone:        "0525856352",
	// }
	// employeeList[2] = employee{
	// 	employeeId:   "103",
	// 	employeeName: "pranee",
	// 	phone:        "0258963258",
	// }
	// fmt.Println("employee =", employeeList)

	employeeList := []employee{}

	employee1 := employee{
		employeeId:   "101",
		employeeName: "Pradoo",
		phone:        "0258963258",
	}

	employee2 := employee{
		employeeId:   "102",
		employeeName: "Prayed",
		phone:        "0258963258",
	}

	employee3 := employee{
		employeeId:   "103",
		employeeName: "PraNee",
		phone:        "0258963258",
	}

	employeeList = append(employeeList, employee1)
	employeeList = append(employeeList, employee2)
	employeeList = append(employeeList, employee3)
	fmt.Println("employee =", employeeList)
}
