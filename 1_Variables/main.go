package main

import (
	"fmt"
	"strconv"
)

// the main is entry point for all go program because compilers looks for entry point while executing go programs
// A program can have only 1 main fucntion, bc you can have only main entrypoint.

// variable declarations

// variables holds two things a memory Locations and a Value

// Package level variable :- (camelCase)
var value_3 int = 80

// Global variable :- (PascalCase)
var Value_2 int = 40

func main() {
	fmt.Println("hello")

	// 1st method

	var a int   // declarations
	a = 10		// initilizations	
	fmt.Println(a)

	// 2nd method

	var b int = 20    // declarations and initilizations, explicit
	fmt.Println(b)

	var c = 30		// less verbose , without datatype
	fmt.Println(c)

	// 3rd method 

	d := "hello"	// shorthand, use only in functions
	fmt.Println(d)

	// different ways 

	var e, f, g = 1, 2, 3   //Multiple Variable Declaration
	fmt.Println("value of E: ",e,"\nvalue of F: ",f,"\nvalue of G: ",g)

	var (
		characters  string
		integers int
		boolean bool
	)

	fmt.Println(characters,"\n",integers,"\n",boolean)

	//NOTE - Variables-Scopes
	// local variable :- inside the functions

	var value int = 20
	fmt.Println(value)
	
	// Accessing.

	companyName := "Mahity-Systems"
	location := "Sanpada"
	hours := 1.5
	reachable := true

	fmt.Printf(
		"name of the company: %s\nlocation of the company: %s\nHours to reach: %v\nis it reachable: %v\n",companyName,location,hours,reachable)

	// Type Checking.
	
	fmt.Printf(
		"name of the company: %s,%T\nlocation of the company: %s,%T\nHours to reach: %v,%T\nis reachable: %v,%T\n",companyName,companyName,location,location,hours,hours,reachable,reachable)

	// Type Conversion

	var x int8 = 20
	fmt.Printf("%v,%T\n",x,x)

	var u int16 = int16(x)
	fmt.Printf("%v,%T\n",u,u)

	var y int = 65
	var str string = strconv.Itoa(y) // converts into string
	var str1 string = string(y) // converts into asci codes


	fmt.Printf("%v,%T\n",str,str)
	fmt.Printf("%v,%T\n",str1,str1)

}
