package main

import "fmt"

type employee struct {
	employee_id   string
	employee_name string
	phone         string
}

func main() {
	employee1 := employee{
		employee_id:   "101",
		employee_name: "Pradoo",
		phone:         "09000000",
	}

	fmt.Println("Employee1 = ", employee1)
}
