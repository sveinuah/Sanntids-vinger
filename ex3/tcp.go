package main

import (
	"fmt"
	"log"
	"net"
	"strings"
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
	port := "34933"

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

	conn, err := net.DialTCP("tcp", localIP, sAddr)
	CheckError(err)

	defer conn.Close()

	fmt.Println(conn.RemoteAddr().String())

}
