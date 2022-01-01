package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("distribuidos", "bui"))
	fmt.Println(strings.Count("distribuidos", "i"))
	fmt.Println(strings.HasPrefix("distribuidos", "di"))
	fmt.Println(strings.HasSuffix("distribuidos", "os"))
	fmt.Println(strings.Index("distribuidos", "bu"))
	fmt.Println(strings.Join([]string{
		"Sistemas",
		"Distribuidos",
		"CUCEI",
	}, "-"))
	fmt.Println(strings.Repeat("distribuidos", 3))
	fmt.Println(strings.Replace("aaabbb", "a", "b", 3))
	fmt.Println(strings.Split("Mi mama me mima", " "))
	tokens := strings.Split("1, show", ",")
	fmt.Println(tokens[2])
	fmt.Println(strings.ToLower("ALEJANDRO"))
	fmt.Println(strings.ToUpper("alejandro"))
	
	//Serializacion
	b := []byte("test")
	fmt.Println(b)

	s := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(s)
}