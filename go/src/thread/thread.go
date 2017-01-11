package main

import (
	. "fmt"
	"runtime"
	"time"
)

func increment(x *int){
	for i := 0; i < 1000000; i++ {
		(*x)++
	}
}

func decrement(x *int){
	for i := 0; i < 1000000; i++ {
		(*x)--
	}
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	i := 0

	go increment(&i)
	go decrement(&i)

	time.Sleep(100*time.Millisecond)
	Println("i = ",i)

}