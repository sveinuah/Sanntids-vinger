package main

import (
	. "fmt"
	"runtime"
)

func increment(accessChan chan int, doneChan chan int, it *int) {
	for i := 0; i < 1000000; i++ {
		myTurn := <- accessChan //If the access variable is in the channel continue, if not wait.
		(*it)++
		accessChan <- myTurn //return access variable
	}
	doneChan <- 1 //send dummy value to say that the routine is done
}

func decrement(accessChan chan int, doneChan chan int, it *int) {
	for i := 0; i < 1000000; i++ {
		myTurn := <- accessChan //If the access variable is in the channel continue, if not wait.
		(*it)--
		accessChan <- myTurn //return access variable
	}
	doneChan <- 1 //send dummy value to say that the routine is done
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	i := 0
	accessChan := make(chan int,1)
	accessChan <- 1

	doneChan := make(chan int)

	go decrement(accessChan, doneChan, &i)
	go increment(accessChan, doneChan, &i)

	<-doneChan
	<-doneChan

	Println("i = ", i)

}
