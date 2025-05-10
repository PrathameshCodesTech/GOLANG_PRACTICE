package main

import "fmt"

type Employee struct{
	// name of attributes and datatypes
	
	Employee_id int
	Employee_name string
	Employee_mobile []string
}


func main() {
	var emp Employee = Employee{
		Employee_id : 12,
		Employee_name : "aman",
		Employee_mobile : []string{"1234567","9876543"},
	}

	employee_2 := Employee{Employee_id : 11,
					Employee_name : "prathamesh",
					Employee_mobile : []string{"1234567","9876543"},}

	fmt.Println(emp.Employee_id)
	fmt.Println(emp.Employee_mobile)
	fmt.Println(emp.Employee_name)

	fmt.Println(employee_2.Employee_name)
	
	emp2 := emp

	// fmt.Println(&emp)
	// fmt.Println(&emp2)

	fmt.Println(&emp == &emp2) // False ,call by pointers

	emp2.Employee_name = "prathamesh"
	// fmt.Println(emp2)
	// fmt.Println(emp)

	emp3 := &emp
	fmt.Println(&emp == emp3) // True ,call by references

	fmt.Println(emp3.Employee_name)
	emp3.Employee_name = "khushi"
	fmt.Println(emp3)
	fmt.Println(emp)

}


