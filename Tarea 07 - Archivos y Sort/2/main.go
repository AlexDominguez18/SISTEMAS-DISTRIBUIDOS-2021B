package main

import (
	"fmt"
	"os"
	"bufio"
	"sort"
)

//Constants
const (
	ADD_PROCESS = 1
	ASCENDING_ORDER = 2
	DESCENDING_ORDER = 3
	EXIT = 4
)

//Struct
type Process struct {
	Id uint64
	Priority int64
	Time uint64
	Status string
}

//Methods
func (p *Process) Show() {
	fmt.Printf("ID = %d, Priority = %d, Time = %d, Status = %s\n", 
				p.Id, p.Priority, p.Time, p.Status)
}

//OrderByPriority
type ByPriority []Process

func (a ByPriority) Len() int           { return len(a) }
func (a ByPriority) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPriority) Less(i, j int) bool { return a[i].Priority < a[j].Priority }

func showMenu() {
	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1. Add process.")
	fmt.Println("2. Ascending Order.")
	fmt.Println("3. Descending Order.")
	fmt.Println("4. Exit.\n")
	fmt.Print("Option: ")
}

func readProcesses(processes *[]Process) {
	var n uint64
	var id uint64 

	fmt.Print("Number of processes: ")
	fmt.Scan(&n)

	if len(*processes) == 0 {
		*processes = nil
	}

	for i := 0; uint64(i) < n; i++ {
		*processes = append(*processes, readProccess(id))
		id ++
	}
}

func readProccess(id uint64) Process {
	var process Process
	scanner := bufio.NewScanner(os.Stdin)
	process.Id = id
	fmt.Println("---------")
	fmt.Print("Priority: ")
	fmt.Scan(&process.Priority)
	fmt.Print("Time: ")
	fmt.Scan(&process.Time)
	fmt.Print("Status: ")
	scanner.Scan()
	process.Status = scanner.Text()
	return process
}

func printProcesses(processes []Process) {
	for i, process := range processes {
		fmt.Printf("%d.- ", i + 1)
		process.Show()
	}
}

func ascendingOrder(processes []Process){
	if len(processes) == 0 {
		fmt.Println("No processes")
		return
	}
	fmt.Println("Before Order:")
	printProcesses(processes)
	sort.Sort(ByPriority(processes))
	fmt.Println("Ascending Order: ")
	printProcesses(processes)
}

func descendingOrder(processes []Process){
	if len(processes) == 0 {
		fmt.Println("No processes")
		return
	}
	fmt.Println("Before Order:")
	printProcesses(processes)
	sort.Sort(sort.Reverse(ByPriority(processes)))
	fmt.Println("Descending Order: ")
	printProcesses(processes)
}

func main() {
	var processes []Process
	
	for isRunning := true; isRunning; {
		var option uint64
		showMenu()
		fmt.Scan(&option)
		fmt.Print("\033[H\033[2J")
		switch option {
		case ADD_PROCESS:
			readProcesses(&processes)
		case ASCENDING_ORDER:
			ascendingOrder(processes)
		case DESCENDING_ORDER:
			descendingOrder(processes)
		case EXIT:
			isRunning = false
		default:
			fmt.Println("Invalid option!")
		}

		if isRunning {
			fmt.Print("Press ENTER to continue...")
			fmt.Scanln()
			fmt.Print("\033[H\033[2J")
		}
	}
}