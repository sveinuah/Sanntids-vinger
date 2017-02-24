package main

import (
	"bcast"
	"fmt"
	"os/exec"
	"time"
)

func main() {
	backupChan := make(chan int, 1)
	rxbackup := make(chan int, 1)

	var id = "backup"
	var num = 0

	go bcast.Receiver(20014, rxbackup)

	active := false
	switch active {
	case false:
		fmt.Println("Hei jeg er", id)

		timer := time.NewTimer(1000 * time.Millisecond)

		for !active {
			select {
			case num = <-rxbackup:
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(1000 * time.Millisecond)
			case <-timer.C:
				active = true
			default:
			}
		}

		fallthrough

	case true:
		id = "master"
		spawnBackup()
		fmt.Println("Jeg er", id)
		mainCounter(backupChan, num)
	}

}

func spawnBackup() {
	term := exec.Command("gnome-terminal", "gnome-terminal", "-e", "go run Sanntids-vinger/ex6/src/backup/backup.go")
	term.Start()

	//startUp := "go run backup.go -id=backup"

	//term.Write([]byte(startUp))
}

func mainCounter(backupChan chan int, num int) {
	go bcast.Transmitter(20014, backupChan)
	printTime := time.Tick(1 * time.Second)
	backupTime := time.Tick(10 * time.Millisecond)
	for {
		select {
		case <-printTime:
			num++
			fmt.Println(num)
		case <-backupTime:
			backupChan <- num
		default:
		}
	}
}
