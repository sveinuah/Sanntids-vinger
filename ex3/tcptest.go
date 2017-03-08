package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

type UnitType struct {
	IP   string
	Port string
}

func main() {

	dataChanTx := make(chan []byte, 1)
	dataChanRx := make(chan []byte, 1)
	connChan := make(chan *net.TCPConn, 1)
	terminateChan := make(chan bool, 1)

	Target := UnitType{"129.241.187.43", "33546"}
	Client := UnitType{"129.241.187.142", "20014"}

	fmt.Println("Initiating TCP Connection... ")
	connChan <- InitiateTCPCon(Target, Client)

	fmt.Print("Starting Communication...")
	go CommunicateOnTCP(connChan, dataChanTx, dataChanRx, terminateChan)
	fmt.Print("Success!")

	go func() {
		received := make([]byte, 1024)
		for {
			select {
			case received <- dataChanRx:
				fmt.Println(string(received))
			default:
			}
		}
	}()

	go func() {
		for {
			dataChanTx <- []byte("Testing Testing, Hallo?")
			time.Sleep(2 * time.Second)
		}
	}()

	time.Sleep(20 * time.Second)
	terminateChan <- true
	conn := <-connChan
	conn.Close()
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}

func InitiateTCPCon(Target UnitType, Client UnitType) (TCPconn *net.TCPConn) {

	tempAddress := Target.IP + ":" + Target.Port
	targetAddress, _ := net.ResolveTCPAddr("tcp", tempAddress)
	clientAddress := Client.IP + ":" + Client.Port
	localAddress, _ := net.ResolveTCPAddr("tcp", clientAddress)
	TCPconn, err := net.DialTCP("tcp", nil, targetAddress)
	CheckError(err)

	conn.Write([]byte("Connect to:" + clientAddress + "\x00"))

	ln, err := net.ListenTCP("tcp", localAddress)
	CheckError(err)

	TCPconn, err = ln.Accept()
	CheckError(err)

	return TCPconn
}

func CommunicateOnTCP(connChan chan *net.TCPConn, dataChanTx chan []byte, dataChanRx chan []byte, terminateChan chan bool) {

	buf := make([]byte, 1024)
	var conn *net.TCPConn
	for {
		conn <- connChan
		if <-terminateChan {
			connChan <- conn
			break
		}
		select {
		case data := <-dataChanTx:
			conn.Write(data)
			fallthrough
		default:
			_, err := conn.Read(buf)
			dataChanRx <- buf
		}
		connChan <- conn
	}
}
