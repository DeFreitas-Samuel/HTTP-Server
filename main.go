package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	fmt.Fprintf(w, "Hola Mundo!")

}
func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	fmt.Fprintf(w, "POST successful")
	message := r.FormValue("message")

	fmt.Fprintf(w, "Message = %s\n", message)
}

func main() {

	fmt.Printf("Starting server at port 8080\n")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/messages", helloWorldFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
