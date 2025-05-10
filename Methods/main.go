package main

import "fmt"

type Calculator struct {
	Num1, Num2 int
}

func (c Calculator) addition() int {
	return c.Num1 + c.Num2
}

func (c Calculator) multiplication() int {
	return c.Num1 * c.Num2
}

func main() {
	v := Calculator{12, 12}
	fmt.Println(v.addition())
	fmt.Println(v.multiplication())
}