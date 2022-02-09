package main

import (
	"fmt"
	"sync"
	"time"
)

//strcuture
type singleTon struct {
	connectionString string
}
var lock = &sync.Mutex{}
var instance *singleTon

func getInstance() *singleTon {
	lock.Lock()
	defer lock.Unlock()
	if instance == nil {
		instance = &singleTon{connectionString: "Kiran"}
	}
	return instance
}

func main() {
	 go func ()  {
		 for i:=0;i<10000000;i++{
			 time.Sleep(time.Microsecond*30)
			 fmt.Println(getInstance(),"-",i)
		 }
	 }()

	 go func ()  {
			fmt.Println(*getInstance()) 
	 }()
	 fmt.Scanln()
}
