package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Whats Your name")
	var name string
	// fmt.Println(name)
	// fmt.Println(&name)  // Address of name

	// fmt.Scan(&name)		// Scan to take inputs in console

	// fmt.Println("your name here ",name) // not comaptible

	reader := bufio.NewReader(os.Stdin)
	name,_ = reader.ReadString('\n')

	fmt.Println(name)




}