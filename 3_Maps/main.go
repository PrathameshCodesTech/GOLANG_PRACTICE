package main

import "fmt"

// MAP = key:value

func main(){

	employee_salary := make(map[string]int) // Declarations, make method
	fmt.Println(employee_salary)

	// Initilizations
	employee_salary = map[string]int{
		"prathamesh" : 0,
		"shubham" : 15000,
	}
	fmt.Println(employee_salary)


	// Declarations + Initilizations
	parent_salary := map[string]int{
		"Papa" : 20000,
		"didi" : 22000,
	}


	fmt.Println(parent_salary)
	//fmt.Println(len(parent_salary))

	//get
	fmt.Println(parent_salary["Papa"])

	//set
	parent_salary["didi"] = 23000
	fmt.Println(parent_salary)

	// update
	parent_salary["mumma"] = 2000
	fmt.Println(parent_salary)


	// delete
	delete(parent_salary,"didi")
	fmt.Println(parent_salary)

	// Call by References

	emp_sal := employee_salary
	fmt.Println(emp_sal)

	emp_sal["prathamesh"] = 3000
	fmt.Println(emp_sal)
	fmt.Println(employee_salary)


	// Traps of MAP

	fmt.Println(employee_salary["aman"]) // it returns 0 but we dosen't know if its salary is 0 or key is not there.
	
	// solution :- Flag, ok

	val,ok := employee_salary["prathamesh"]
	fmt.Println(val,ok)

	// without variable

	_,flag := employee_salary["prathamesh"]
	fmt.Println(flag)

}