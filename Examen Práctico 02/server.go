package main

import (
	"fmt"
	"net"
	"net/rpc"
	"encoding/gob"
	"strconv"
)

const (
	MOVIES_PORT = ":9001"
	GAMES_PORT = ":9002"
	MUSIC_PORT = ":9003"
	MIDDLEWARE_PORT = ":9000"
	MOVIES = 1
	GAMES = 2
	MUSIC = 3
)

type Client struct {
	Id uint64
	Conn net.Conn
}

type ChatRoom struct {
	Topic string
	Port string
	ConnectedClients uint64
}

type Middleware struct {
	ChatRooms []ChatRoom
}

func (this *Middleware) ChatRoomStatus(dummie int, reply *string) error {
	*reply = "CHAT ROOMS\n----------\n"
	for i := 0; i < len(this.ChatRooms); i++ {
		*reply += strconv.Itoa(i + 1) + 
			"-" + this.ChatRooms[i].Topic +
			", People connected: " + strconv.Itoa(int(this.ChatRooms[i].ConnectedClients)) + "\n"
	}
	return nil
}

func (this *Middleware) GetChatRoomStatus(dummie int, reply *string) error {
	*reply += "<table class='table table-bordered'>" + 
			"<thead><tr>" +
			"<th>#</th> " + 
			"<th>Topic</th>" +
			"<th>IP</th>" +
			"<th>People connected</th>" +
			"</tr></thead>"
	for i := 0; i < len(this.ChatRooms); i++ {
		*reply += "<tbody><tr>" +
			"<td>" + strconv.Itoa(i + 1) + "</td>" +
			"<td>" + this.ChatRooms[i].Topic + "</td>" +
			"<td>" + "127.0.0.1:" + strconv.Itoa(9000 + i + 1) + "</td>" +
			"<td>" + strconv.Itoa(int(this.ChatRooms[i].ConnectedClients)) + "</td>" + 
			"</tr><tbody>"
	}
	*reply += "</table>"
	return nil
}

func (this *Middleware) GetChatRoom(option int, reply *string) error {
	switch option {
	case MOVIES:
		*reply = MOVIES_PORT
	case GAMES:
		*reply = GAMES_PORT
	case MUSIC:
		*reply = MUSIC_PORT
	default:
		*reply = "Error: Unknown option"
	}
	return nil
}

func middleware(m *Middleware) {
	rpc.Register(m)
	middleware, err := net.Listen("tcp", ":9000")
	
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	
	for {
		client, err := middleware.Accept()
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		go rpc.ServeConn(client)
	}
}

func server(chatRoom *ChatRoom) {
	var clients []Client = make([]Client, 0)
	var messages []string = make([]string, 0)
	var clientId uint64

	server, err := net.Listen("tcp", chatRoom.Port)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	//Listen connections
	go func(){
		for {
			c, err := server.Accept()
			if err != nil {
				fmt.Println("Error: ", err)
				continue
			}
			client := Client{
				Id: clientId, 
				Conn: c,
			}
			clients = append(clients, client)
			//Handle Client
			go func() {
				for {
					var msg string
					gob.NewDecoder(c).Decode(&msg)
					if msg == "disconnect" {
						c.Close()
						chatRoom.ConnectedClients--
						clients = append(clients[:client.Id], clients[:client.Id+1]...)
						return
					} else {
						messages = append(messages, msg)
						for i := 0; i < len(clients); i++ {
							gob.NewEncoder(clients[i].Conn).Encode(msg)
						}
					}
				}
			}()
			clientId++
			chatRoom.ConnectedClients++
		}
	}()
}

func main() {
	var m Middleware
	
	m.ChatRooms = []ChatRoom{
		ChatRoom{"Movies", MOVIES_PORT, 0},
		ChatRoom{"Games", GAMES_PORT, 0},
		ChatRoom{"Music", MUSIC_PORT, 0},
	}
	go middleware(&m)
	go server(&m.ChatRooms[0])
	go server(&m.ChatRooms[1])
	go server(&m.ChatRooms[2])
	fmt.Println("Running server...")
	fmt.Scanln()
}