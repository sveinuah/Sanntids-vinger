package main

import (
	. "fmt"
	"runtime"
)

var activeGoRoutines[2] int

func increment(varChan chan int, it *int) {
	for i := 0; i < 1000000; i++ {
		myTurn := <- varChan
		(*it)++
		varChan <- myTurn
	}
	activeGoRoutines[0] = 0
}

func decrement(varChan chan int, it *int) {
	for i := 0; i < 1000000; i++ {
		myTurn := <- varChan
		(*it)--
		varChan <- myTurn
	}
	activeGoRoutines[1] = 0
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	i := 0
	varChan := make(chan int,1)
	varChan <- 1

	activeGoRoutines[1] = 1
	go decrement(varChan, &i)
	activeGoRoutines[0] = 1
	go increment(varChan, &i)

	for(activeGoRoutines[0]==1 || activeGoRoutines[1]==1){
	}

	Println("i = ", i)

}
