package main

import (
	. "fmt"
	"runtime"
	"time"
)

func increment(var *int){
	for ( i := 0; i < 1000000; i++){
		(*var)++
	}
}

func decrement(var *int){
	for(i := 0; i < 1000000; i++){
		(*var)++
	}
}

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())

	i := 0

	go increment()
	go decrement()

	time.Sleep(100*time.Millisecond)
	Println("i = %d",i)

}