package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func cargarHtml(file string) string {
	html, _ := ioutil.ReadFile(file)
	return string(html)
}

func root(rest http.ResponseWriter, req *http.Request) {
	//Retornar algo al navegador
	rest.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(rest, cargarHtml("index.html"))
}

func main() {
	http.HandleFunc("/", root)
	fmt.Println("Arrancando el servidor...")
	http.ListenAndServe(":9000", nil)
}