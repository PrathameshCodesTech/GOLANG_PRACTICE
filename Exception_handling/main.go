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

func main1() {

	// div := divide(20, 2)
	// fmt.Println("the result is: ",div)
	
	div,err := divide(20, 0)
	fmt.Println("the result is: ",div)
	if err != nil{
		fmt.Println("Error :",err.Error())
	}

	
}


func multiply(num1,num2 int)(int,error){
	if num2 == 0{
		return 1,fmt.Errorf("number cannot be zero")
	}
	return num1 * num2,nil
}

func main() {
	a,b := multiply(12,2)
	if b != nil{
		fmt.Println(b)
	}
	fmt.Println(a)
	
}