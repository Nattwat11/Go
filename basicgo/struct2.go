package main

import "fmt"

type employee struct {
	employee_name string
	employee_id   string
	phone         string
}

/*
employee_list := [2]employee_list{}

	employee_list[0]  = employee{
		employee_id = "0012",
		employee_name = "DO"
		phone = "0900002103"
	employee_list[1] = employee{
		employee_id = "0013",
		employee_name = "AUM"
		phone  = "0900002310"
	}
	fmt.Prinln("employee = ",employee_list)
	}
*/
func main() {
	employee1 := employee{
		employee_id:   "100",
		employee_name: "Bank",
		phone:         "090000001",
	}

	employee2 := employee{
		employee_id:   "101",
		employee_name: "Pradoo",
		phone:         "09000000",
	}
	employee3 := employee{
		employee_id:   "102",
		employee_name: "Pra",
		phone:         "0900001",
	}
	employee4 := employee{
		employee_id:   "103",
		employee_name: "Pranee",
		phone:         "09000003",
	}
	employee_list := []employee{}
	employee_list = append(employee_list, employee1)
	employee_list = append(employee_list, employee2)
	employee_list = append(employee_list, employee3)
	employee_list = append(employee_list, employee4)
	fmt.Println("employee =", employee_list)
}
