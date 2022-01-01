package main

import (
	"fmt"
	"net"
)

func servidor() {
	s, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := s.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleClient(c)
	}
}

func handleClient(c net.Conn) {
	b := make([]byte, 100)
	_, err := c.Read(b)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	go servidor()
}