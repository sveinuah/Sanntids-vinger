package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func main() {
	sAddr, err := net.ResolveUDPAddr("udp", ":30000")
	CheckError(err)
	fmt.Println(sAddr)

	//lAddr, err := net.ResolveUDPAddr("udp", "localhost:30000")

	//conn, err := net.DialUDP("udp", nil, sAddr)
	//CheckError(err)

	conn, err := net.ListenUDP("udp", sAddr)
	CheckError(err)
	defer conn.Close()

	buf := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buf)
	CheckError(err)
	fmt.Println("Received: ", string(buf[0:n]), "From: ", addr)

	serverAddr, err := net.ResolveUDPAddr("udp", "129.241.187.43:20014")

	conn2, err := net.DialUDP("udp", nil, serverAddr)
	CheckError(err)
	defer conn2.Close()

	quitChan := make(chan int, 2)

	go listen(conn2, quitChan)
	go write(serverAddr, conn2, quitChan)
	/*
		for {
			conn2.WriteTo(msg, serverAddr)
			n, _, _ := conn2.ReadFromUDP(buf)

			fmt.Println("Received:", string(buf[0:n]))

			time.Sleep(2 * time.Second)
		}
	*/

	<-quitChan
	<-quitChan
}

func listen(conn *net.UDPConn, quitChan chan int) {
	for {
		fmt.Println("Tullball")
		buf := make([]byte, 1024)
		n, _, _ := conn.ReadFromUDP(buf)

		fmt.Println("Received:", string(buf[0:n]))
	}
}

func write(serverAddr *net.UDPAddr, conn *net.UDPConn, quitChan chan int) {
	for {
		fmt.Println("SAddr:", serverAddr.String())
		message := "Heo Sveio!"
		msg := make([]byte, 1024)
		copy(msg[:], message)
		conn.WriteTo(msg, serverAddr)
		time.Sleep(2 * time.Second)
	}
}
