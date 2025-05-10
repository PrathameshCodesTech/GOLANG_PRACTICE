package main

import (
	"errors"
	"fmt"
)

func divide(num1, num2 int) (int,error) {
	if num2==0{
		return -1,errors.New("num2 is 0")
	}
	return num1 / num2,nil
}

func main() {

	// div := divide(20, 2)
	// fmt.Println("the result is: ",div)
	
	div,err := divide(20, 0)
	fmt.Println("the result is: ",div)
	if err != nil{
		fmt.Println("Error :",err.Error())
	}

	
}