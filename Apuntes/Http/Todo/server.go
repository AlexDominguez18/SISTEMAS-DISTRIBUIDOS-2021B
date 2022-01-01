package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Tarea struct {
	Nombre string
	Estado string
}

type AdminTareas struct {
	Tareas []Tarea
}

func (tareas *AdminTareas) Agregar(tarea Tarea) {
	tareas.Tareas = append(tareas.Tareas, tarea)
}

var misTareas AdminTareas

func cargarHtml(file string) string {
	html, _ := ioutil.ReadFile(file)
	return string(html)
}

func (tareas *AdminTareas) String() string {
	var html string
	for _, tarea := range tareas.Tareas {
		html += "<tr>" +
		"<td>" + tarea.Nombre + "</td>" +
		"<td>" + tarea.Estado + "</td>" +
		"</tr>"
	}
	return html
}

func form(rest http.ResponseWriter, r *http.Request) {
	rest.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		rest,
		cargarHtml("form.html"),
	)
}
func tareas(rest http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(rest, "ParseForm() error %v", err)
				return
			}
			tarea := Tarea{
				Nombre: r.FormValue("tarea"),
				Estado: r.FormValue("estado"),
			}
			misTareas.Agregar(tarea)
			rest.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				rest, 
				cargarHtml("respuesta.html"),
				tarea.Nombre,
			)
		case "GET":
			rest.Header().Set(
				"Content-type",
				"text/html",
			)
			fmt.Fprintf(
				rest,
				cargarHtml("tareas.html"),
				misTareas.String(),
			)
	}
}

func main() {
	http.HandleFunc("/form", form)
	http.HandleFunc("/tareas", tareas)
	fmt.Println("Arrancando el servidor...")
	http.ListenAndServe(":9000", nil)
}