package main

import "fmt"

func main1(){

	// Primitive DT
	a := 5
	b := 10
	fmt.Println(a == b)  // false
	fmt.Println(a != b)  // true
	fmt.Println(a < b)   // true


	// Arrays are comparable if their elements are comparable.
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1 == arr2) // true — same length and values


	// In Go, nil means "no memory is allocated" for a variable that’s a reference type — like:
	// slices, maps, pointers, channels,functions,	interfaces

	
	// var a []int         // nil slice
	// b := []int{}        // empty slice

	// fmt.Println(a == nil) // true — no memory
	// fmt.Println(b == nil) // false — memory allocated for zero-length array

	// fmt.Println(len(a), cap(a)) // 0 0
	// fmt.Println(len(b), cap(b)) // 0 0




	// Slices can only be compared to nil.
	s1 := []int{1, 2, 3}
	// s2 := []int{1, 2, 3}

	var s2 []int
	
	fmt.Println(s2==nil) // TRUE
	// fmt.Println(s1 == s2) //  compile error: invalid operation
	fmt.Println(s1 == nil)   // false — only comparison allowed



	// Like slices, maps can only be compared to nil.
	m1 := map[string]int{"a": 1}
	// m2 := map[string]int{"a": 1}
	// fmt.Println(m1 == m2) 					//  compile error
	fmt.Println(m1 == nil)   					// false — allowed

	
	
	// 5. Structs (comparable if all fields are comparable)
	//  all fields are comparable, struct is comparable.
	// If the struct contains a slice, map, or function, it becomes non-comparable:

	type Person struct {
		Name string
		Age  int
	}
	
	p1 := Person{"Aman", 25}
	p2 := Person{"Aman", 25}
	fmt.Println(p1 == p2)			 // true

}



func main2() {
	a := 10
	b := 3

	fmt.Println("Addition:", a + b)       // 13
	fmt.Println("Subtraction:", a - b)    // 7
	fmt.Println("Multiplication:", a * b) // 30
	fmt.Println("Division:", a / b)       // 3 (integer division)
	fmt.Println("Modulus:", a % b)        // 1

	// Go is strict about types:

	// a := 10
	// b := 3.0
	// fmt.Println(a + b) //  compile error (mismatched types)

}


func main3() {
	a := 10

	a += 5   // a = a + 5 → 15
	fmt.Println("After += :", a)

	a -= 3   // a = a - 3 → 12
	fmt.Println("After -= :", a)

	a *= 2   // a = a * 2 → 24
	fmt.Println("After *= :", a)

	a /= 4   // a = a / 4 → 6
	fmt.Println("After /= :", a)

	a %= 5   // a = a % 5 → 1
	fmt.Println("After %= :", a)
}


func main() {
	a := true
	b := false

	fmt.Println("a && b:", a && b) // false — both must be true
	fmt.Println("a || b:", a || b) // true — one is true
	fmt.Println("!a:", !a)         // false — inverts true
	fmt.Println("!b:", !b)         // true  — inverts false

	age := 25
	valid := true

	if age > 18 && valid {
		fmt.Println("Eligible")
	} else {
		fmt.Println("Not eligible")
}

}





