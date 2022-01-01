package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"encoding/json"
)

type Admin struct {
	Students map[string]map[string]float64
	Subjects map[string]map[string]float64
}

var admin Admin

func readHtml(fileName string) string {
	html, _ := ioutil.ReadFile(fileName)
	return string(html)
}

func index(rest http.ResponseWriter, request *http.Request) {
	rest.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		rest,
		readHtml("index.html"),
	)
}

func create(rest http.ResponseWriter, request *http.Request) {
	rest.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		rest,
		readHtml("create.html"),
	)
}

func response(rest http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "POST":
			if err := request.ParseForm(); err != nil {
				fmt.Fprintf(rest, "ParseForm() error: %v", err)
				return
			}
			student := request.FormValue("student")
			subject := request.FormValue("subject")
			grade, _ := strconv.ParseFloat(request.FormValue("grade"), 64)
			if _, err := admin.Students[student]; err == false {
				admin.Students[student] = make(map[string]float64)
			}
			if _, err := admin.Subjects[subject]; err == false {
				admin.Subjects[subject] = make(map[string]float64)
			}
			admin.Students[student][subject] = grade
			admin.Subjects[subject][student] = grade
			rest.Header().Set("Content-Type", "text/html")
			j, _ := json.Marshal(admin.Students)
			response := string(j)
			j, _ = json.Marshal(admin.Subjects)
			response += "\n" + string(j)
			fmt.Println(response)
			fmt.Fprintf(rest, readHtml("response.html"), string(response))
	}
}

func searchStudent(rest http.ResponseWriter, request *http.Request) {
	rest.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		rest,
		readHtml("searchStudent.html"),
	)
}

func showStudent(rest http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "GET":
			student := request.FormValue("student")
			if _, err := admin.Students[student]; err == false {
				return
			}
			var gradeSummary float64
			for _, grade := range admin.Students[student] {
				gradeSummary += grade
			}
			average := gradeSummary/float64(len(admin.Students[student]))
			rest.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(rest, readHtml("showStudent.html"), student, average)
	}
}

func showGeneralAverage(rest http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "GET":
			var averagesSummary float64
			var average float64
			for name, _ := range admin.Students {
				var gradeSummary float64
				for _, grade := range admin.Students[name] {
					gradeSummary += grade
				}
				averagesSummary += gradeSummary/float64(len(admin.Students[name]))
			}
			average = averagesSummary/float64(len(admin.Students))
			rest.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				rest,
				readHtml("showGeneralAverage.html"),
				average,
			)
	}
}

func searchSubject(rest http.ResponseWriter, request *http.Request) {
	rest.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		rest,
		readHtml("searchSubject.html"),
	)
}

func showSubject(rest http.ResponseWriter, request *http.Request) {
	switch request.Method {
		case "GET":
			subject := request.FormValue("subject")
			if _, err := admin.Subjects[subject]; err == false {
				return
			}
			var gradeSummary float64
			var average float64
			for _, grade := range admin.Subjects[subject] {
				gradeSummary += grade
			}
			average = gradeSummary / float64(len(admin.Subjects[subject]))
			rest.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				rest,
				readHtml("showSubject.html"),
				subject,
				average,
			)
	}
}

func main() {
	//Admin initialization
	admin.Students = make(map[string]map[string]float64)
	admin.Subjects = make(map[string]map[string]float64)
	//Routes
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/response", response)
	http.HandleFunc("/searchStudent", searchStudent)
	http.HandleFunc("/showStudent", showStudent)
	http.HandleFunc("/showGeneralAverage", showGeneralAverage)
	http.HandleFunc("/searchSubject", searchSubject)
	http.HandleFunc("/showSubject", showSubject)
	//Message
	fmt.Println("Running server...")
	//Server
	http.ListenAndServe(":8080", nil)
}