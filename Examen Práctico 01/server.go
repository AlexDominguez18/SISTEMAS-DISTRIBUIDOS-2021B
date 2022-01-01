package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
	"strconv"
	"./util"
)

const (
	MESSAGES_FILES = 1
	BACKUP = 2
	EXIT_SERVER = 3
)

type File struct {
	Name string
	Bytes []byte
}

type Client struct {
	Id int
	Conn net.Conn
}

func closeServer(clients *[]Client) {
	for _, c := range *clients {
		os.RemoveAll(strconv.Itoa(c.Id))
	}
	os.RemoveAll("server")
}

func showMenu() {
	util.ClearScreen()
	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1. Show messages/files.")
	fmt.Println("2. Backup messages/files.")
	fmt.Println("3. Exit.")
	fmt.Print("Option: ")
}

func handleClient(c net.Conn, clients *[]Client, messages *[]string, id int) {
	err := os.Mkdir(strconv.Itoa(id), 0755)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	for {
		var msg string
		gob.NewDecoder(c).Decode(&msg)
		switch msg {
		case "connect":
			break
		case "disconnect":
			c.Close()
			*clients = append((*clients)[:id], (*clients)[id+1:]...)
			os.RemoveAll(strconv.Itoa(id))
			return
		case "file":
			var f File
			gob.NewDecoder(c).Decode(&f)
			*messages = append((*messages), f.Name)
			file, err := os.Create("server/" + f.Name)
			if err != nil {
				fmt.Println("Error: ", err)
				break
			}
			file.Write(f.Bytes)
			file.Close()
			for i := 0; i < len(*clients); i++ {
				file, err = os.Create(strconv.Itoa((*clients)[i].Id) + "/" + f.Name)
				if err != nil {
					fmt.Println("Error: ", err)
					break
				}
				file.Write(f.Bytes)
				file.Close()
				gob.NewEncoder((*clients)[i].Conn).Encode(f.Name)
			}
		default:
			*messages = append(*messages, msg)
			for i := 0; i < len(*clients); i++ {
				gob.NewEncoder((*clients)[i].Conn).Encode(msg)
			}
		}
	}
}

func main() {
	var isRunning bool = true
	var option int
	var clients []Client = make([]Client, 0)
	var messages []string = make([]string, 0)
	
	err := os.Mkdir("server", 0755)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	s, err := net.Listen("tcp", ":9999")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	go func() {
		var clientId int
		for {
			c, err := s.Accept()
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			clients = append(clients, Client{Id: clientId, Conn: c})
			go handleClient(c, &clients, &messages, clientId)
			clientId++
		}
	}()

	for ; isRunning ; {
		showMenu()
		fmt.Scan(&option)
		util.ClearScreen()
		switch option {
		case MESSAGES_FILES:
			fmt.Println("Messages/Files:")
			for _, msg := range messages {
				fmt.Println("-> ", msg)
			}
		case BACKUP:
			if len(messages) == 0 {
				fmt.Println("Nothing to backup!")
				break
			}
			currentTime := time.Now()
			file, err := os.Create(currentTime.String() + "-backup.txt")
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
			defer file.Close()
			for _, msg := range messages {
				file.WriteString(msg + "\n")
			}
		case EXIT_SERVER:
			closeServer(&clients)
			isRunning = false
		default:
			fmt.Println("Unknown option")
		}
		
		if isRunning {
			util.Pause()
		}
	}
}