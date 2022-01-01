package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"./util"
)

const (
	MESSAGE = 1
	FILE = 2
	SHOW = 3
	EXIT_CLIENT = 4
)

type File struct {
	Name string
	Bytes []byte
	Extension string
}

func showMenu() {
	util.ClearScreen()
	fmt.Println("MENU")
	fmt.Println("---------")
	fmt.Println("1. Send message.")
	fmt.Println("2. Send file.")
	fmt.Println("3. Show messages.")
	fmt.Println("4. Exit.")
	fmt.Print("Option: ")
}

func main() {
	var isRunning bool = true
	var option int
	var messages []string = make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	
	c, err := net.Dial("tcp", ":9999")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	//Connection
	gob.NewEncoder(c).Encode("connect")

	go func() {
		for {
			var msg string
			gob.NewDecoder(c).Decode(&msg)
			messages = append(messages, msg)
		}
	}()

	for ; isRunning; {
		showMenu()
		fmt.Scan(&option)
		switch option {
		case MESSAGE:
			var msg string
			fmt.Print(": ")
			scanner.Scan()
			msg = scanner.Text()
			gob.NewEncoder(c).Encode(msg)
		case FILE:
			gob.NewEncoder(c).Encode("file")
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("File name: ")
			scanner.Scan()
			fileName := scanner.Text()
			file, err := os.Open(fileName)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			defer file.Close()
			stat, err := file.Stat()
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			var f File
			f.Name = fileName
			f.Bytes = make([]byte, stat.Size())
			count, err := file.Read(f.Bytes)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			str := string(f.Bytes)
			fmt.Println(count,str)
			gob.NewEncoder(c).Encode(f)
		case SHOW:
			util.ClearScreen()
			fmt.Println("CHAT ROOM\n---------")
			for _, msg := range messages {
				fmt.Println("-", msg)
			}
		case EXIT_CLIENT:
			isRunning = false
			gob.NewEncoder(c).Encode("disconnect")
		default:
			fmt.Println("Unknown option")
		}
		
		if isRunning {
			util.Pause()
		}
	}
}