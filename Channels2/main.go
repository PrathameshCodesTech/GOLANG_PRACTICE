package main

import "fmt"

func display(channel chan int) {
	fmt.Println("Go Routines")
	for i:=1;i<=3;i++{
		channel<-i
	}
	close(channel)
}

func main(){
	mychannel:=make(chan int)
	go display(mychannel)
	for{
		num,status:= <- mychannel
		if status == false{ // if channel is closed
			break
		}
		fmt.Println(num)
	}
	fmt.Println("Main Goroutine")
}