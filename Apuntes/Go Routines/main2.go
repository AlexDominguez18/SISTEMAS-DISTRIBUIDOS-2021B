package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func () {
		for {
			c1 <- "Canal 1"
			time.Sleep(time.Second * 2)
		}
	}() //Para poder ejecutar la funcion
	
	go func () {
		for {
			c2 <- "Canal 2"
			time.Sleep(time.Second * 1)
		}
	}() //Para poder ejecutar la funcion

	go func () {
		for {
			select {
			case msg := <- c1:
				fmt.Println(msg)
			case msg := <- c2:
				fmt.Println(msg)
			}
		}
	}()

	fmt.Scanln()
}