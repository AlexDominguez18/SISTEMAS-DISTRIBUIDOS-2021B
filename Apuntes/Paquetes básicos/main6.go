package main

import (
	"fmt"
	"sort"
)

type Persona struct{
	nombre string
	edad uint64
}

type ByNombre []Persona

func (a ByNombre) Len() int           { return len(a) }
func (a ByNombre) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByNombre) Less(i, j int) bool { return a[i].nombre < a[j].nombre }

type ByEdad []Persona

func (a ByEdad) Len() int           { return len(a) }
func (a ByEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByEdad) Less(i, j int) bool { return a[i].edad < a[j].edad }

func main() {
	personas := []Persona{
		Persona{nombre: "Alex", edad: 21},
		Persona{nombre: "Denis", edad: 54},
		Persona{nombre: "Juan", edad: 42},
	}

	sort.Sort(ByNombre(personas))
	fmt.Println(personas)	

	sort.Sort(ByEdad(personas))
	fmt.Println(personas)
}