package main

import (
	"fmt"
)

func main() {
	// arrays:

	// var arr [5]int = [5]int{1, 2, 3, 5}  // creation of arrays
	// var arr_1 = [...]int{1,23,4,5,9}     // dyanmic arrays
	
	// Slice Arrays

	//NOTE -  1st way

	var num [5]int							// fixed 
	var num_2 []int							// variable sized

	fmt.Printf("Fixed size %T \n",num)
	fmt.Printf("Variable size %T \n",num_2)

	fmt.Println("Look at console",num)
	fmt.Println("Look at console",num_2)

	num[0] = 10								// we can do this in arrays
	num[1] = 20
	
	// num_2[0] = 20							// We can't do this in slice as range is not mentioned

	// WHAT WE CAN DO IS WE CAN APPEND THE VALUE.

	num_2 = append(num_2, 10,20,30)
	fmt.Println("Look at console",num)
	fmt.Println("Look at console",num_2)

	//NOTE - 2nd way

	var arr []int = []int{1,2,34,4,5,6,7}
	// fmt.Printf("%v,%T",arr,arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

	// Using For Loop we can append

	for i:=0; i<5;i++{
		arr = append(arr, i)
	}
	fmt.Println(arr)


	// another syntax
	var array []int = make([]int,2,4)
	fmt.Println(len(array))
	fmt.Println(cap(array))


	// in make method, we can indexes value, since range is given.

	array[1] = 4

	// Append with another variable 
	arr_2 := append(array,3)
	fmt.Println(array)
	fmt.Println(arr_2)


	// append with same
	array = append(array, 45)
	fmt.Println(array)
	fmt.Println(arr_2)


	slyce_arrays := make([]int, 4, 16) 					// Create a slice of len=4, cap=16
	slyce_arrays = append(slyce_arrays, 1, 23, 45, 54) 			// Append values

	slyce_arrays_2 := slyce_arrays // Copy the slice (shares same underlying array)
	slyce_arrays_2 = append(slyce_arrays_2, 34)

	fmt.Println(slyce_arrays)
	fmt.Println(slyce_arrays_2)

	slyce_arrays = append(slyce_arrays,100)  			// after appending og slice, copy slice values gets truncated bcs of og length exceeds

	fmt.Println(slyce_arrays)
	fmt.Println(slyce_arrays_2)

}


func main2(){
	var slice_arrs []int = []int{10, 20, 30, 40}
	fmt.Println("Original slice:", slice_arrs)

	var slice_arrs2 *[]int = &slice_arrs
	fmt.Println("Pointer:", slice_arrs2)
	fmt.Println("Dereferenced value:", *slice_arrs2)

	*slice_arrs2 = append(*slice_arrs2, 50, 60, 70)

	fmt.Println("Updated slice_arrs:", slice_arrs)
	fmt.Println("Updated *slice_arrs2:", *slice_arrs2)
}

func main3(){
	original := make([]int, 4, 5)  // Length 4, capacity is 5
	original[0], original[1], original[2], original[3] = 1, 2, 3, 4

	copyRef := original          // same backing array

	fmt.Println(original)
	fmt.Println(copyRef)


	copyRef[0] = 11				// both slice will get triggered

	fmt.Println(original)        
	fmt.Println(copyRef)

	copyRef = append(copyRef, 99) // triggers new array! even though, there is capacity

	// copyRef[0] = 111
	fmt.Println("original:", original)
	fmt.Println("copyRef:", copyRef)   
	
}