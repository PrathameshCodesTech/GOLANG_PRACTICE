package main

import "fmt"

func main(){
	fmt.Println()
	obj := summation(12,12)
	fmt.Println("Summation of two number is",obj)
	obj1:=concat()
	fmt.Println(obj1)
	obj2 := multiplication(12,12)
	fmt.Println("multiply of two number is",obj2)

	obj3,obj4 := mul_add(12,12)
	// _,obj4 := mul_add(12,12)  // to get single value
	fmt.Println("multiply of two number is",obj3,"Summation of two number is",obj4)

	obj5 := add(12,12,12,12,2,2,12,12)
	fmt.Println("Summation of numbers is",obj5)

}

func summation(num1 int, num2 int) int{
	return (num1 + num2)
}

func concat()string{
	return "hello" + "world"
}

func multiplication(num1,num2 int)int{
	return (num1*num2)
}

// return multiple value

func mul_add(num1,num2 int)(int,int){
	x := num1 + num2
	y := num1 * num2
	return x,y
}

// Varidatic function in golang / variable size

func add(nums ...int)int{
	s :=0
	for _,value := range nums{
		s += value
	}
	return s
}

// Anonymous Functions

 