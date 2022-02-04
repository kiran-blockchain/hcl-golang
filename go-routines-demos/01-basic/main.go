package main

import (
	"fmt"
	"net/http"
	"runtime"

	//"runtime"
	"time"
)

func routineCheck(from string) {
	start := time.Now()
	for i := 0; i < 100000000; i++ {
		//fmt.Println(from, ":", i)
	}
	elapsed := time.Since(start)
	fmt.Println("Routine Completed in ",elapsed)
}
func webCheck(_link string){
_,err := http.Get(_link)
if err!=nil {
	fmt.Println(_link,"might be down")
	return
}
fmt.Println(_link,"is up")
}
func main() {
	// regular function call
	//routineCheck("Direct Call")
	runtime.GOMAXPROCS(1)
	start := time.Now()
	links := []string{
		"https://google.com",
		"https://fb.com",
		"https://golang.org",
	}
	for _,link := range links{
		fmt.Println("Calling ",link);
		go webCheck(link)
	}
	// this is the way to call any function as a routine.
	// for i := 0; i < 16; i++ {
	// 	go routineCheck("Routine" + string(i))
	// }
	elapsed := time.Since(start)
	time.Sleep(time.Second * 5)
	fmt.Println("Done")
	fmt.Println("Time lapsed for execution", elapsed)
}
