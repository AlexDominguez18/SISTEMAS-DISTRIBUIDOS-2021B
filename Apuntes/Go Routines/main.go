package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Millisecond * 1000)
	}
}

func main() {
	//El channel se debe crear con make
	var c chan string = make(chan string)
	//Procesos concurrentes
	go pinger(c)
	go ponger(c)
	go printer(c)

	fmt.Scanln()
}