package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Tarea struct {
	ID uint64 `json:"id"`
	Nombre string `json:"nombre"`
	Estado string `json:"estado"`
}

var admin map[uint64]Tarea

func Get() ([]byte, error) {
	jsonData, err := json.MarshalIndent(admin, "", "\t")
	if err != nil {
		return jsonData, nil
	}
	return jsonData, err
}

func GetID(id uint64) ([]byte, error) {
	jsonData := []byte(`{}`)
	tarea, ok := admin[id]
	if ok == false {
		return jsonData, nil
	}
	jsonData, err := json.MarshalIndent(tarea, "", "\t")
	if err != nil {
		return jsonData, err
	}
	return jsonData, nil
}

func Add(tarea Tarea) ([]byte) {
	jsonData := []byte(`{"code":"ok"}`)
	admin[tarea.ID] = tarea
	return jsonData
}

func Delete(id uint64) ([]byte) {
	_, ok := admin[id]
	if ok == false {
		return []byte(`{"code":"noexiste"}`)
	}
	delete(admin, id)
	return []byte(`{"code":"ok"}`)
}

func Update(id uint64, tarea Tarea) ([]byte) {
	_, ok := admin[id]
	if ok == false {
		return []byte(`{"code":"noexiste"}`)
	}
	admin[id] = tarea
	return []byte(`{"code":"ok"}`)
}

//Rutas

func tarea(res http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			fmt.Println("GET")
			res_json, err := Get()
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set(
				"Content-Type", 
				"application/json",
			)
			res.Write(res_json)
		case "POST":
			fmt.Println("POST")
			var tarea Tarea
			err := json.NewDecoder(r.Body).Decode(&tarea)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res_json := Add(tarea)
			res.Header().Set(
				"Content-Type",
				"application/json",
			)
			res.Write(res_json)
	}
}

func tarea_id(res http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/tarea/"), 10, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(r.Method)
	switch r.Method {
		case "GET":
			res_json, err := GetID(id)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
			}
			res.Header().Set(
				"Content-Type",
				"application/json",
			)
			res.Write(res_json)
		case "DELETE":
			res_json := Delete(id)
			res.Header().Set(
				"Content-Type",
				"application/json",
			)
			res.Write(res_json)
		case "PUT":
			var tarea Tarea
			err := json.NewDecoder(r.Body).Decode(&tarea)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res_json := Update(id, tarea)
			res.Header().Set(
				"Content-Type",
				"application/json",
			)
			res.Write(res_json)
	}
}

func main() {
	admin = make(map[uint64]Tarea)

	http.HandleFunc("/tarea", tarea)
	http.HandleFunc("/tarea/", tarea_id)
	fmt.Println("Corriendo RESTful API...")
	http.ListenAndServe("localhost:9000", nil)
}