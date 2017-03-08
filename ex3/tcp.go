package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func getServerMsg() *net.UDPAddr {
	sAddr, err := net.ResolveUDPAddr("udp", ":30000")
	CheckError(err)

	conn, err := net.ListenUDP("udp", sAddr)
	CheckError(err)

	buf := make([]byte, 1024)

	n, addr, err := conn.ReadFromUDP(buf)
	CheckError(err)
	fmt.Println("Received: ", string(buf[0:n]), "From: ", addr)

	conn.Close()
	return addr
}

func LocalIP() string {
	var localIP string

	conn, err := net.DialTCP("tcp4", nil, &net.TCPAddr{IP: []byte{8, 8, 8, 8}, Port: 53})
	CheckError(err)
	defer conn.Close()
	localIP = strings.Split(conn.LocalAddr().String(), ":")[0]

	return localIP
}

func main() {
	port := "33546"

	tempAddr := getServerMsg()
	splitAddr := strings.Split(tempAddr.String(), ":")
	serverAddr := splitAddr[0] + ":" + port
	fmt.Println(serverAddr)
	sAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	CheckError(err)

	lIP := LocalIP()
	lIP = lIP + ":" + "20014"
	fmt.Println(lIP)
	localIP, err := net.ResolveTCPAddr("tcp", lIP)
	CheckError(err)

	conn, err := net.DialTCP("tcp", nil, sAddr)
	CheckError(err)

	defer conn.Close()

	fmt.Println(conn.RemoteAddr().String())

	//message := "Da tester vi TCP!"
	//iter := 0

	//Ask for connection between server and client
	waitChan := make(chan bool, 10)

	//waitChan <- false
	reqCon := "Connect to: " + lIP + "\x00"
	fmt.Println(reqCon)

	conn.Write([]byte(reqCon))

	listener, err := net.ListenTCP("tcp", localIP)
	CheckError(err)
	TCPconn, err := listener.Accept()
	fmt.Println("Listen tcp")

	CheckError(err)

	buf := make([]byte, 1024)
	reqCon = "heihei\x00"
	go func() {
		for {
			_, err = TCPconn.Read(buf)

			CheckError(err)
			TCPconn.Write([]byte(reqCon))

			fmt.Println(string(buf))
			time.Sleep(2 * time.Second)

		}
	}()
	<-waitChan
}
