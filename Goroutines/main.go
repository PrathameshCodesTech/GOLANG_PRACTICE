// Concurrency - Goroutines
// Goroutines are lightweight threads managed by the Go runtime
// -> they let you run functions concurrently.
// -> resides in same address
// -> Main Goroutines should wait for while untill all the child goroutines returned back.
// -> Goroutines are separte lightweight thread being managed by Go Container

// Basic Syntax
//  add go before a function call to run it concurrently!

// go functionName()

package main

import (
	"fmt"
	"time"
	// "time"
)

func printMessage(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, i)
		// time.Sleep(time.Millisecond * 500)
	}
}


func main1() {		// Main Goroutines 

	go printMessage("From Goroutine")
	time.Sleep(1 * time.Second)
	fmt.Println("Runs in main thread")
	


	// printMessage("From Main")         // runs in main thread
	// go printMessage("From Goroutine") // runs in background -> fork a new goroutine to execute the function call print message 
}


// The main function exits immediately, so goroutines might not finish unless you wait, usually using:

// time.Sleep() (hacky for testing),sync.WaitGroup (best practice), Channels


func display(name string){
	for i:=1; i<=5;i++{
		fmt.Println(i,"=",name)
	}
}


func main(){
	go display("Ambasoft java is goated youtube channel")
	time.Sleep(1 * time.Second)
	fmt.Println("Waiting state")
	fmt.Println("Main Goroutines")
}


