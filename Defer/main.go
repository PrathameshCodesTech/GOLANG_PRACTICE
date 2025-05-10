package main


import "fmt"

func main() {
	fmt.Println("Start")

	defer fmt.Println("Deferred: This runs last")

	fmt.Println("Middle")

	// Function returns here — now the deferred call runs
}


//  Defer runs in LIFO (Last-In, First-Out) order.
func example() { 
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	fmt.Println("Inside function")
}
