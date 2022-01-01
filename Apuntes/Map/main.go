package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	// m := make(map[string]int)
	// m["Programacion"] = 100

	m := make(map[string]map[string]int)
	m["Primero"] = make(map[string]int)
	m["Primero"]["Programacion"] = 100
	m["Primero"]["Algoritmia"] = 90
	m["Segundo"] = make(map[string]int)
	m["Segundo"]["Bases de datos"] = 100

	j, err := json.Marshal(m)

	if err != nil {
		return
	}

	fmt.Println(string(j))
}