package main

import (
	"fmt"
	"net"
	"time"
	"encoding/gob"
)

type Process struct {
	Id uint64
	Count uint64
}

func (p *Process) RunProcess() {
	for {
		(*p).Count = (*p).Count + 1
		time.Sleep(time.Millisecond * 500)
	}
}

func handleClient(c net.Conn, processes *[]Process) {
	var msg string
	gob.NewDecoder(c).Decode(&msg)

	if msg == "connect" {
		gob.NewEncoder(c).Encode((*processes)[0])
		*processes = append((*processes)[:0], (*processes)[1:]...)
		c.Close()
	} else {
		var p Process
		gob.NewDecoder(c).Decode(&p)
		*processes = append(*processes, p)
		c.Close()
	}
}

func server() {
	var processes []Process = make([]Process, 0)

	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	for i := uint64(0); i < 5; i++ {
		var process = Process{ 
			Id: i, 
			Count: 0,
		}
		processes = append(processes, process)
	}
	for i := uint64(0); i < uint64(len(processes)); i++ {
		go processes[i].RunProcess()
	}

	go func() {
		for {
			for _, p := range processes {
				fmt.Printf("%d : %d\n",p.Id, p.Count)
			}
			fmt.Println("----------------")
			time.Sleep(time.Millisecond * 501)
		}
	}()

	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		go handleClient(c, &processes)
	}
}

func main() {
	gob.Register(Process{})
	go server()

	var input string
	fmt.Scan(&input)
}