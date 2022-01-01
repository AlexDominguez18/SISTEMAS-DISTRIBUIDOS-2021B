package main

import (
	"fmt"
	"time"
	"strings"
	"strconv"
)

const (
	ADD_PROCESS    = 1
	SHOW_PROCESS   = 2
	DELETE_PROCESS = 3
	EXIT           = 4
	SHOW = "show"
	NOT_SHOW = "not show"
	DELETE = "delete"
	DELIMITER = ","
)

type Process struct {
	id uint64
}

func (p *Process) start(c chan string) {
	go RunProcess(p.id, c)
}

func (p *Process) stop(c chan string) {
	c <- strconv.Itoa(int(p.id)) + DELIMITER + DELETE
}

func RunProcess(id uint64, c chan string) {
	i := uint64(0)
	isShown := false
	for {
		select {
		case msg := <-c:
			tokens := strings.Split(msg, ",")
			processId, err := strconv.Atoi(tokens[0])
			if err != nil {
				return
			}
			if  uint64(processId) == id {
				if tokens[1] == SHOW {
					isShown = true
				} else if tokens[1] == NOT_SHOW {
					isShown = false
				} else {
					return
				}
			} else {
				c <- msg
			}
		default:
			if isShown {
				fmt.Printf("ID %d: %d\n", id, i)
			}
			i = i + 1
			time.Sleep(time.Millisecond * 500)	
		}
	}
}

func getProccesById(id uint64, processes []Process) int64 {
	for index, p := range processes {
		if p.id == id {
			return int64(index) 
		}
	}
	return -1
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func pauseScreen() {
	var input string
	fmt.Print("Press ENTER to continue...")
	fmt.Scanln(&input)
}

func menu() {
	clearScreen()
	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1. Add process.")
	fmt.Println("2. Show processes.")
	fmt.Println("3. Delete process.")
	fmt.Println("4. Exit.")
	fmt.Print("\nOption: ")
}

func main() {
	var isRunning bool = true
	var option int
	var id uint64
	var processes []Process = make([]Process, 0)
	var channel chan string = make(chan string)

	for isRunning {
		menu()
		fmt.Scan(&option)
		clearScreen()
		switch option {
		case ADD_PROCESS:
			process := Process{id: id}
			processes = append(processes, process)
			process.start(channel)
			id++
			fmt.Println("New process started!")
		case SHOW_PROCESS:
			clearScreen()
			for _, p := range processes {
				channel <- strconv.Itoa(int(p.id)) + DELIMITER + SHOW
			}
			fmt.Scan(&option)
			if option == 0 {
				for _, p := range processes {
					channel <- strconv.Itoa(int(p.id)) + DELIMITER + NOT_SHOW
				}
			}
		case DELETE_PROCESS:
			var idToDelete uint64
			fmt.Print("Process's Id to delete: ")
			fmt.Scan(&idToDelete)
			index := getProccesById(idToDelete, processes)
			if index >= 0 {
				processes[index].stop(channel)
				fmt.Printf("Deleted process %d!\n", processes[index].id)
				processes = append(processes[:index], processes[index+1:]...)
			} else {
				fmt.Println("Invalid process's Id!")
			}
		case EXIT:
			isRunning = false
		default:
			fmt.Println("Invalid option!")
		}

		if isRunning {
			pauseScreen()
		}
	}

}
