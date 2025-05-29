package main

import (
	"fmt"
)

type employee struct {
	name                 string
	days                 []int
	total_available_days int
	total_used_days      int
}

func main() {
	var employees []employee
	var total_employees int
	fmt.Print("Enter the number of employees to schedule: ")
	_, _ = fmt.Scanf("%d", &total_employees)
	fmt.Println(total_employees)
	employees = make([]employee, total_employees, total_employees)
	for i := range total_employees {
		fmt.Printf("Enter the employee name n. %d: ", i+1)
		var name string
		_, _ = fmt.Scanf("%s", &name)
		fmt.Println(name)
		employees[i] = employee{
			name,
			[]int{0, 0, 0, 0, 0},
			0,
			0,
		}
	}
	fmt.Println(employees)
}
