package main

import (
	"fmt"
	"bufio"
	"os"
	"net"
	"net/rpc"
	"encoding/gob"
	"./util"
)

const (
	MOVIES_CHATROOM = "Movies Chat Room"
	GAMES_CHATROOM = "Games Chat Room"
	MUSIC_CHATROOM = "Music Chat Room"
	SEND_MESSAGE = 1
	SHOW_MESSAGES = 2
	EXIT = 3
)

func showMenu(topic string) {
	util.ClearScreen()
	fmt.Println(topic)
	fmt.Println("---------")
	fmt.Println("1. Send message.")
	fmt.Println("2. Show messages.")
	fmt.Println("3. Exit.")
	fmt.Print("Option: ")
}

func main() {
	//Variables
	var scanner = bufio.NewScanner(os.Stdin)
	var messages []string = make([]string, 0)
	var reply string
	var option int64

	//RPC Connection
	rpcClient, err := rpc.Dial("tcp", "127.0.0.1:9000")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	//Middleware connection
	err = rpcClient.Call("Middleware.ChatRoomStatus", 0, &reply)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(reply)
	fmt.Print("Choose a chat room: ")
	fmt.Scan(&option)
	err = rpcClient.Call("Middleware.GetChatRoom", option, &reply)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	rpcClient.Close()
	//Chatroom connection
	client, err := net.Dial("tcp", reply)
	if err != nil{
		fmt.Println("Error: ", err)
		return
	}
	//Messages
	go func() {
		for {
			var msg string
			gob.NewDecoder(client).Decode(&msg)
			messages = append(messages, msg)
		}
	}()
	var topic string
	switch reply {
	case ":9001":
		topic = MOVIES_CHATROOM
	case ":9002":
		topic = GAMES_CHATROOM
	case ":9003":
		topic = MUSIC_CHATROOM
	}
	for {
		showMenu(topic)
		fmt.Scan(&option)
		switch option {
		case SEND_MESSAGE:
			var msg string
			fmt.Print("Type: ")
			scanner.Scan()
			msg = scanner.Text()
			gob.NewEncoder(client).Encode(msg)
		case SHOW_MESSAGES:
			util.ClearScreen()
			fmt.Println(topic, " - Messages")
			for _, msg := range messages {
				fmt.Println("-", msg)
			}
		case EXIT:
			gob.NewEncoder(client).Encode("disconnect")
			return
		}
		util.Pause()
	}
}