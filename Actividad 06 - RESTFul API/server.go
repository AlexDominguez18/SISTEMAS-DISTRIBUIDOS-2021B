package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"strconv"
)

type Record struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Subject string `json:"subject"`
	Grade float64 `json:"grade"`
}

type Student struct {
	ID uint64 `json:"id"`
	Name string `json:"name"`
	Subjects map[string]float64 `json:"subjects"`
}

var admin map[uint64]Student

func Get() ([]byte, error) {
	jsonData, err := json.MarshalIndent(admin, "", "\t")
	if err != nil {
		return jsonData, nil
	}
	return jsonData, err
}

func GetID(id uint64) ([]byte, error) {
	jsonData := []byte(`{}`)
	student, ok := admin[id]
	if ok == false {
		return jsonData, nil
	}
	jsonData, err := json.MarshalIndent(student, "", "\t")
	if err != nil {
		return jsonData, err
	}
	return jsonData, nil
}

func Add(r Record) ([]byte) {
	jsonData := []byte(`{"code":"ok"}`)
	if student, ok := admin[r.ID]; ok == false {
		student.Subjects = make(map[string]float64)
		student.ID = r.ID
		student.Name = r.Name
		admin[r.ID] = student
	}
	admin[r.ID].Subjects[r.Subject] = r.Grade
	return jsonData
}

func Update(id uint64, student Student) ([]byte) {
	_, ok := admin[id]
	if ok == false {
		return []byte(`{"code":"Null"}`)
	}
	admin[id] = student
	return []byte(`{"code":"ok"}`)
}

func Delete(id uint64) ([]byte) {
	_, ok := admin[id]
	if ok == false {
		return []byte(`{"code":"Null"}`)
	}
	delete(admin, id)
	return []byte(`{"code":"ok"}`)
}

//Routes

func student(res http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	switch r.Method {
		case "GET":
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
			var record Record
			err := json.NewDecoder(r.Body).Decode(&record)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res_json := Add(record)
			res.Header().Set(
				"Content-Type",
				"application/json",
			)
			res.Write(res_json)
	}
}

func student_id(res http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	id, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/student/"), 10, 64)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
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
		case "PUT":
			var student Student
			err := json.NewDecoder(r.Body).Decode(&student)
			fmt.Println(student)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res_json := Update(id, student)
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
	}
}

func main() {
	admin = make(map[uint64]Student)
	//Routes
	http.HandleFunc("/student", student)
	http.HandleFunc("/student/", student_id)
	fmt.Println("Running server...")
	http.ListenAndServe(":9000", nil)
}