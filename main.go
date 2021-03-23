package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Text struct {
	Mensajes []string
}

var Messages = []string{"Hello", "World"}

func helloWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")

	switch r.Method {

	case "GET":
		text := Text{Messages}
		js, err := json.Marshal(text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		postMessage := r.FormValue("postMessage")
		fmt.Fprintf(w, "postMessage = %s\n", postMessage)
		Messages = append(Messages, postMessage)

	default:
		fmt.Fprintf(w, "This URL only suports GET and POST methods")
	}

}
func helloWorldFunc2(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "PUT":
		URL := (r.URL.Path)
		split := strings.Split(URL, "/")
		valueToEdit, err := strconv.ParseInt(split[2], 10, 64)
		if err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
		}
		fmt.Printf(split[2])
		putMessage := r.FormValue("putMessage")
		fmt.Fprintf(w, "putMessage = %s\n", putMessage)
		Messages[valueToEdit] = putMessage

	case "DELETE":
		fmt.Fprintf(w, "No implemented yet")

	default:
		fmt.Fprintf(w, "This URL only suports PUT and DELETE methods")
	}
}

func main() {

	fmt.Printf("Starting server at port 8080\n")
	http.HandleFunc("/messages", helloWorldFunc)
	http.HandleFunc("/messages/", helloWorldFunc2)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
