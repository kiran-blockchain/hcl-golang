package main

import "fmt"

func routine(ch chan int) {
	fmt.Println(100 + <-ch)
}
func main() {
	channel_demo := make(chan int)
	// channel is accepting int

	go routine(channel_demo)
	channel_demo <- 234 
	//sending the data to the channel
	
	fmt.Println("Value of channel ", channel_demo)
	fmt.Printf("Type of channel %T ", channel_demo)
}
