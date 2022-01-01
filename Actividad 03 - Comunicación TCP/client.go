package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
	"os/signal"
	"os"
)

type Process struct {
	Id    uint64
	Count uint64
}

func (p *Process) RunProcess(c chan string) {
	for {
		select {
		case msg := <- c:
			if msg == "finish" {
				return
			}
		default:
			fmt.Printf("%d : %d\n", p.Id, p.Count)
			(*p).Count = (*p).Count + 1
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func client() {
	var channel chan string = make(chan string)
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	msg := "connect"
	gob.NewEncoder(c).Encode(msg)
	var p Process
	for {
		gob.NewDecoder(c).Decode(&p)
		if p != (Process{}) {
			fmt.Println("Connected: ", p)
			break
		}
		c.Close()
	}

	go p.RunProcess(channel)

	key := make(chan os.Signal, 1)
	signal.Notify(key, os.Interrupt)
	go func(){
		for sig := range key {
			fmt.Println("Signal: ", sig)
			channel <- "finish"
			c2, err := net.Dial("tcp", ":9999")
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			msg := "disconnect"
			gob.NewEncoder(c2).Encode(&msg)
			gob.NewEncoder(c2).Encode(p)
			fmt.Println("Finished client!")
			return
		}
	}()
}

func main() {
	gob.Register(Process{})
	go client()

	var input string
	fmt.Scan(&input)
}
