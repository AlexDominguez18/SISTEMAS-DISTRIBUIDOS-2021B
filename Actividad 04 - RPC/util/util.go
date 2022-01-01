package util

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func Pause() {
	var input string
	fmt.Println("Press ENTER to continue...")
	fmt.Scanln(&input)
}