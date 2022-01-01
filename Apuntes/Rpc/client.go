package main

import (
	"fmt"
	"net/rpc"
)

func client() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var op int64
	for {
		fmt.Println("1) Hello")
		fmt.Println("2) Negate")
		fmt.Println("3) Min")
		fmt.Println("0) Exit")
		fmt.Scanln(&op)

		switch op {
		case 1:
			var name string
			fmt.Print("Name: ")
			fmt.Scanln(&name)
			var result string
			err = c.Call("Server.Hello", name, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Hello=", result)
			}
		case 2:
			var number int64
			fmt.Print("Number: ")
			fmt.Scanln(&number)
			var result int64
			err = c.Call("Server.Negate", number, &result)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server.Negate", number, "=", result)
			}
		case 3:
			slice := []int64{0, 0, 0, 0}
			for i := 0; i < len(slice); i++ {
				fmt.Print(i, "= ")
				fmt.Scanln(&slice[i])
			}
			var result int64
			err = c.Call("Server.Min", slice, &result)
			if err != nil {
				fmt.Println(err)
			} else {h
				fmt.Println("Server.Min", slice, "=", result)
			}
		case 0:
			return
		}
	}
}

func main() {
	client()
}