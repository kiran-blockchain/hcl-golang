package main

import "fmt"

func myfunc(channel chan string){
	for i:=0;i<7;i++{
		channel <- "Go Lang is Awesome "
	}
	fmt.Println("Lengh of the channel is", len(channel))
	close(channel)
}
func main (){
	c := make(chan string,5)
	go myfunc(c)
	//my func is acting as the data publisher
	fmt.Println("Capacity of the channel is", cap(c))
	counter:=0
	for {
		res, ok  := <-c
		counter++
		if !ok {
			fmt.Println("Channel is closed", ok)
			break
		}
		fmt.Println("Channel is open and data is", res,ok)
	}
}