package main

import (
	"./network/bcast"
	"fmt"
	"time"
)

type HelloMsg struct {
	Message string
	Iter    int
}

func main() {
	helloTx := make(chan HelloMsg)
	helloRX := make(chan HelloMsg)

	go bcast.Transmitter(12000, helloTx)
	go bcast.Receiver(12000, helloRX)

	go func() {
		helloMsg := HelloMsg{"Echo? .. Echo? ... Echo?", 0}
		for {
			helloMsg.Iter++
			helloTx <- helloMsg
			time.Sleep (2 * time.Second)
		}
	}()

	fmt.Println("Program started:")
	for {
		select {
		case a := <-helloRX:
			fmt.Println("Received:", a)
		}
	} 	
}
