package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	backup := make(chan int)
	startChan := make(chan bool)
	go mainCounter(backup, startChan, os.Stdout, 0)

}

func mainCounter(backup chan int, startChan chan bool, stdout *os.File, num int) {
	go spawnBackup(backup, startChan)
	os.Stdout = stdout
	fmt.Println("Starting count")
	for {
		num++
		fmt.Println(num)
	}

}

func spawnBackup(backup chan int, startChan chan bool) {
	cmd := exec.Command("start", "cmd")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(cmd.Stdout, "Hei! Jeg er backup!")

	var num int
	running := true
	for running == true {
		select {
		case num = <-backup:
		case running = <-startChan:
		}
	}
	mainCounter(backup, startChan, cmd.Stdout, num)
}
