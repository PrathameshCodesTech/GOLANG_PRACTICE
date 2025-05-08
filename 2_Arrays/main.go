package main

import (
	"fmt"
)

func main() {
	// syntax of arrays 

	var arr [5]int = [5]int{1,0,2,34,5}
	fmt.Println(arr)

	for i := 0;i< len(arr);i++{
		fmt.Println(i,arr[i])
	}

	// dyanmic
	var arr1 = [...]int{1,2,23,3,4,4,6,7,8,9}
	// arr1 := [...]int{1,2,23,3,4,4,6,7,8,9}
	fmt.Println(arr1)

	// methods

	// fmt.Println(len(arr))
	// fmt.Println(len(arr1))

	// updates/reassign at particular posistion

	arr1[9] = 10
	// fmt.Println(arr1)
	// fmt.Println(arr1[5])

	// call by value

	arr2 := arr
	// fmt.Println(arr2)
	// fmt.Println(arr)


	arr2[0] = 20

	// fmt.Println(arr2)
	// fmt.Println(arr)

	// fmt.Println(arr1[:5])

	// Matrix of arrays 

	var arrays[3][4]int = [3][4]int{{1,2,3,6},{2,3,1},{7,6,5}}
	fmt.Println(arrays)
	fmt.Println(arrays[2][3])
	
	for i:=0;i<len(arrays);i++{
		for j:=0;j<len(arrays[i]);j++{
			fmt.Println(arrays[i][j]," ")
		}
		fmt.Println()
	}
}