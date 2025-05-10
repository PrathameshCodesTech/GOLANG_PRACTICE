// Concurrency - Channels

// -> Are Pipes through which we can pass data to Goroutines.
// -> sends and receive data from a cahnnel

package main

import (
	"fmt"
	// "time"
)

func main1(){
	// special varaible that is use to exchange informations among goroutines
	
	var mychannel chan int 		// Declarations of vaiable mychannel as an int channel
	fmt.Println(mychannel)

	if mychannel == nil{
		fmt.Println("mychannel is not initiaized and it will return nill")
	}

	mychannel = make(chan int) // creating and int channel
	fmt.Println(mychannel)

	if mychannel != nil{
		fmt.Println("mychannel is initiaized")
	}

	// data:= <- mychannel //Read from the channel  -> arrows goes away from channel 
	// mychannel <-  data	// writing to channels	-> goes towards channel

}



// Whenever Value being read from channel. Go runtime will block the execution untill a value gets written to the channels 
// -> by Default channels are bi-directional. Send and Receive.
// -> INSTEAD OF SLEEP, we re using channels it will replace time.sleep
// -> if there is read operation from a channel inside a goroutines and there should be a proper a write operation to supply data to channel, otherwise, go compiler will raise a panic of deadlock

func display2(channel chan int){
	fmt.Println("display goroutine, a single thread")
	channel<-10 		// writes value 10

}

func main2(){
	mychannel := make(chan int)			// creating an integer channel

	go display2(mychannel)
	// <-mychannel							// Reading data from channel from my channels  ->Block the execution untill a write happens to the channel
	// -> if blocking is not there we should use sleep.
	
	num:=<-mychannel	
	fmt.Println(num)

	fmt.Println(" Main Goroutines ")	// Terminates main Goroutines


}



// Uni-Directional Channels


// Sends only channel
//	-> Write only channels, We cannot read from channels, we can just send/write data into it.

// Receive only Channels 
	// -> we cant write a value to channles we can jus receive or read data from it

// Conerting Bidirectianl channles to Unidirectional
	// -> it is possible to convert channel either a send only or receive-only

	
					// Write 			// Read			
func display(channel chan<- int,channel2 <- chan int){
	fmt.Println("Display Goroutines")
	channel <- 20
	//channel2 := <- 40 // Test this
	
}

func main(){
	// mychannel := make(chan<- int)   // Sends only channel
	mychannel := make(chan int)   // channels as bidirectinal
	// go display(mychannel)
	// num := <- mychannel  	-> We cannot read from channels.
	// fmt.Println("supplied data ", num)

	mychannel2 := make(<- chan int)
	go display(mychannel,mychannel2)

	fmt.Println("Main Goroutines")
}