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

func getServerMsg() {
	sAddr, err := net.ResolveUDPAddr("udp", ":30000")
	CheckError(err)

	conn, err := net.ListenUDP("udp", sAddr)
	CheckError(err)

	buf := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buf)
	CheckError(err)
	fmt.Println("Received: ", string(buf[0:n]), "From: ", addr)

	conn.Close()
}

func main() {

	getServerMsg()

	quitChan := make(chan int, 2)

	listenAddr, _ := net.ResolveUDPAddr("udp", ":20014")
	conn2, err := net.ListenUDP("udp", listenAddr)
	CheckError(err)
	defer conn2.Close()

	go listen(conn2, quitChan)
	go write(conn2, quitChan)

	<-quitChan
	<-quitChan
}

func listen(conn *net.UDPConn, quitChan chan int) {
	for {
		buf := make([]byte, 1024)
		n, _, _ := conn.ReadFromUDP(buf)

		fmt.Println("Received:", string(buf[0:n]))
	}
}

func write(conn *net.UDPConn, quitChan chan int) {
	serverAddr, err := net.ResolveUDPAddr("udp", "129.241.187.43:20014")
	CheckError(err)

	for {
		_, err := conn.WriteToUDP([]byte("yo!"), serverAddr)
		CheckError(err)
		time.Sleep(2 * time.Second)
	}
}
