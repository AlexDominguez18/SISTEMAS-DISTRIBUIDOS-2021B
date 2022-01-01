package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/rpc"
	// "strconv"
	// "encoding/json"
)

func readHtml(fileName string) string {
	html, _ := ioutil.ReadFile(fileName)
	return string(html)
}

func index(rest http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rest.Header().Set("Content-Type", "text/html")
		admin, err := rpc.Dial("tcp", "127.0.0.1:9000")
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		var reply string
		err = admin.Call("Middleware.GetChatRoomStatus", 0, &reply)
		if err == nil {
			fmt.Fprintf(
				rest,
				readHtml("index.html"),
				reply,
			)	
		} else {
			fmt.Println("Error: ", err)
		}
	}
}

func main() {
	http.HandleFunc("/", index)	
	http.ListenAndServe(":8080", nil)
}