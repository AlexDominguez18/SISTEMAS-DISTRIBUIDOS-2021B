package main

import (
	"fmt"
	"net"
	"net/rpc"
	"encoding/json"
)

type Args struct{
	Subject string 
	Student string
	Grade float64
}

type Server struct { 
	Subjects map[string]map[string]float64
	Students map[string]map[string]float64
}

func (this *Server) AddSubjectGrades(args Args, reply *string)  error {
	if _, err := this.Subjects[args.Subject]; err == false {
		this.Subjects[args.Subject] = make(map[string]float64)
	}
	if _, err := this.Students[args.Student]; err == false {
		this.Students[args.Student] = make(map[string]float64)
	}
	this.Subjects[args.Subject][args.Student] = args.Grade
	this.Students[args.Student][args.Subject] = args.Grade
	j, err := json.Marshal(this.Subjects)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	*reply = string(j)
	j, err = json.Marshal(this.Students)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	*reply += "\n" + string(j)
	return nil
}

func (this *Server) ShowStudentAverage(student string, reply *float64) error {
	if _, err := this.Students[student]; err {
		var summary float64
		for _, grade := range this.Students[student] {
			summary += grade	
		}
		*reply = summary / float64(len(this.Students[student]))
	} else {
		*reply = -1
	}
	return nil
}

func (this *Server) ShowGeneralAverage(zero int, reply *float64) error {
	if len(this.Students) > zero {
		var generalAverage float64
		var studentAverage float64
		for name, _ := range this.Students {
			studentAverage = 0
			for _, grade := range this.Students[name] {
				studentAverage += grade
			}
			generalAverage += (studentAverage / float64(len(this.Students[name])))
		}
		generalAverage = generalAverage / float64(len(this.Students))
		*reply = generalAverage
	}
	return nil
}

func (this *Server) ShowSubjectAverage(subject string, reply *float64) error {
	if _, err := this.Subjects[subject]; err {
		var summary float64
		for _, grade := range this.Subjects[subject] {
			summary += grade
		}
		*reply = summary / float64(len(this.Subjects[subject]))
	}
	return nil
}

func server(server *Server) {
	rpc.Register(server)
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	var s Server
	s.Subjects = make(map[string]map[string]float64)
	s.Students = make(map[string]map[string]float64)

	go server(&s)

	var input string
	fmt.Scanln(&input)
}