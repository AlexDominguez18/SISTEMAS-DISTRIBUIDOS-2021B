package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

//Constants
const (
	ASCENDING_FILE = "ascending.txt"
	DESCENDING_FILE = "descending.txt"
	ENTER_STRINGS = 1
	ASCENDING_ORDER = 2
	DESCENDING_ORDER = 3
	EXIT = 4
)

func showMenu(){
	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1. Enter strings.")
	fmt.Println("2. Ascending Order.")
	fmt.Println("3. Descending Order.")
	fmt.Println("4. Exit.\n")
	fmt.Print("Option: ")
}

func readStrings(strs *[]string) {
	var n int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("How many strings do you will write? : ")
	fmt.Scan(&n)

	if len(*strs) > 0 {
		*strs = nil
	}

	for i := 0; i < n; i++ {
		fmt.Print(": ")
		scanner.Scan()
		*strs = append(*strs, scanner.Text())
	}
}

func ascendingOrder(strs []string){
	file, err := os.Create(ASCENDING_FILE)

	if len(strs) == 0 {
		fmt.Println("There is not strings to order")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sort.Strings(strs)
	for _, str := range strs {
		file.WriteString(str + "\n")
	}
}

func descendingOrder(strs []string) {
	file, err := os.Create(DESCENDING_FILE)

	if len(strs) == 0{
		fmt.Println("There is not strings to order.")
		return
	}

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sort.Sort(sort.Reverse(sort.StringSlice(strs)))
	for _, str := range strs {
		file.WriteString(str + "\n")
	}
}

func main() {
	strs := []string{}
	var option uint64

	for isRunning := true; isRunning; {
		showMenu()
		fmt.Scan(&option)
		fmt.Print("\033[H\033[2J")
		switch option {
		case ENTER_STRINGS:
			readStrings(&strs)
		case ASCENDING_ORDER:
			ascendingOrder(strs)
		case DESCENDING_ORDER:
			descendingOrder(strs)
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