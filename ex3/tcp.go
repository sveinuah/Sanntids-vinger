package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
		conn, _ := net.Dial("tcp","127.0.0.2:12000")
		for {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Text to send: ")
			text, _ := reader.ReadString('\n')
			fmt.Fprintf(conn, text + '\n')
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Print("Message from server: "+message)	
		}
}
