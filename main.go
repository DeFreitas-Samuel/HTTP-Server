package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Text struct {
	Mensajes []string
}

func helloWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	var Messages = []string{"Hello", "World"}
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

	case "PUT":
		fmt.Fprintf(w, "No implemented yet")
	case "DELETE":
		fmt.Fprintf(w, "No implemented yet")
	default:
		fmt.Fprintf(w, "This Web Server only suports GET, POST, PUT and DELETE methods")
	}

}

func main() {

	fmt.Printf("Starting server at port 8080\n")
	http.HandleFunc("/messages", helloWorldFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
