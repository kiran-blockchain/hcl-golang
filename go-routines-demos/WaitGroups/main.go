package main

import (
	"fmt"
	"sync"
	"time"
)

func routine( i int, wg *sync.WaitGroup){
 fmt.Println("Started Rounte", i)
 time.Sleep(time.Second*2)
 fmt.Printf("Routine %d ended \n",i)
 wg.Done()
}

func main (){
	noOfRoutines:=3
	var wg sync.WaitGroup
	for i:=0;i<noOfRoutines;i++{
		wg.Add(1)
		go routine(i,&wg)
	}
	wg.Wait();
	fmt.Println("All routines are finished")
}