package main

import (
	"fmt"
	"reflect"
)

type Employee struct {
	name     string
	salary   int
	position string
}

type Company struct {
	company_name string
	employees    []Employee
}

func init_func() Company {
	emp1 := Employee{
		name:     "Haroon Ayaz",
		salary:   1000,
		position: "Lead Machine Learning Engineer",
	}

	emp2 := Employee{
		name:     "Vro",
		salary:   1000,
		position: "Lead Robotics Engineer",
	}

	emps := []Employee{emp1, emp2}

	company := Company{
		company_name: "Stark Industries",
		employees:    emps,
	}
	return company
}

func main() {
	company := init_func()
	size := len(company.employees)
	fmt.Printf("got something init?\n")
	fmt.Printf("%s\n", reflect.TypeOf(company).String())
	fmt.Printf("===========================\n")
	fmt.Printf("Company: %s\n", company.company_name)
	fmt.Println("   --- Total Employees: ", size)
	emp := company.employees
	for i := 0; i < size; i++ {
		fmt.Println("---- Employee No. :", i+1)
		fmt.Println(" Emp Name : ", emp[i].name)
		fmt.Println(" Emp Salary : ", emp[i].salary)
		fmt.Println(" Emp Position: ", emp[i].position)
	}
	fmt.Println("===========================")
}
