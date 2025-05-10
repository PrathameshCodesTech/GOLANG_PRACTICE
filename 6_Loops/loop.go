package main

import "fmt"




func main() {
	//  Classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println("hello",i)
	}

	// While-style loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}




	// using slice

	nums := []int{10, 20, 30}
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// using map

	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// using arrays

	var n [5]int = [5]int{10,20,30,40,50}
	for a,b := range n{
		fmt.Println(a,b)
	}

	m2 := map[string]int{"a":12,"c":24}
	for _,b:=range m2{
		fmt.Println(b)
	}
}