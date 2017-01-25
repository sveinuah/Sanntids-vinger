package main

import (
	//"./Network-go/network/bcast"
	"./Network-go/network/conn"
	"./Network-go/network/localip"
	//"./Network-go/network/peers"
	. "fmt"
	"log"
	"net"
)

type HelloMsg struct {
	Message string
	Iter    int
}

func main() {
	localid, err := localip.LocalIP()
	if err != nil {
		log.Fatal(err)
	}
	Println(localid)
	pConn := conn.DialBroadcastUDP(30000)

	var buf [1024]byte
	n, addr, _ := pConn.ReadFrom(buf[0:])
	Println(n)
	Println(addr)
	var str string
	for i := 0; i < n; i++ {
		str += string(buf[i])
	}
	Printf("%q", str)

	var tAddr net.Addr

	tempAddr := addr.String()
	for i := range buf {
		if string(tempAddr[i]) == ":" {
			break
		}

		tAddr.String() += string(tempAddr[i])
	}
	var tBuf [1024]byte
	pConn.WriteTo(tBuf, addr)
}
