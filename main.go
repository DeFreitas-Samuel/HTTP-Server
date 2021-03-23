package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Messages = []string{"Hello", "World"}

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	fmt.Fprintf(w, "POST successful")
	message := r.FormValue("message")

	fmt.Fprintf(w, "Message = %s\n", message)
}

func helloWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	switch r.Method {
	case "GET":
		js, err := json.Marshal(Messages)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	case "POST":
		//fmt.Fprintf(w, "This Web Server only suports GET, POST, PUT and DELETE methods")
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		postMessage := r.FormValue("postMessage")
		fmt.Fprintf(w, "postMessage = %s\n", postMessage)
		Messages = append(Messages, postMessage)

	default:
		fmt.Fprintf(w, "This Web Server only suports GET, POST, PUT and DELETE methods")
	}

}

func main() {

	fmt.Printf("Starting server at port 8080\n")
	//fileServer := http.FileServer(http.Dir("./static"))
	//http.Handle("/", fileServer)
	//http.HandleFunc("/form", formHandler)

	http.HandleFunc("/messages", helloWorldFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
