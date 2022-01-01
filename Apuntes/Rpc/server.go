package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Server struct {	
}

func (this *Server) Hello(name string, reply *string) error {
	*reply = "Hello " + name
	return nil
}

func (this *Server) Negate(i int64, reply *int64) error {
	*reply = -i
	return nil
}

func (this *Server) Min(s []int64, reply *int64) error {
	if len(s) > 0 {
		min := s[0]
		for i := 1; i < len(s); i++{
			if s[i] < min {
				min = s[i]
			}
		}
		*reply = min
		return nil
	} else {
		return errors.New("Empty slice")
	}
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}