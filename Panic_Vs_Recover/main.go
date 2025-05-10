// What is panic?
// It stops the normal flow of the program.

// It's often used when something unrecoverable happens.

// panic can be recovered using recover() inside a defer block.

package main

import ("fmt"

"os")

func main1() {
	fmt.Println("Start")

	panic("Something went terribly wrong!")

	fmt.Println("This will not be printed") // Never executed
}



func main2() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	panic("Unexpected error!")
	fmt.Println("After panic") // Won’t run
}



func main() {
	// 1. Panic: stops, runs defer, shows stack trace

	defer fmt.Println("This will run if panic happens")
	// panic("Something broke!") // uncomment to test

	// 2. log.Fatal: stops, does NOT run defer, prints message
	// log.Fatal("Fatal error!") // uncomment to test

	// 3. os.Exit: exits with code, no defer, no stack trace
	fmt.Println("Before Exit") // This prints
	os.Exit(1)                 // uncomment to test
	fmt.Println("After Exit")  // This won't print
}

