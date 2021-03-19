package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	fmt.Fprintf(w, "Hello World Hola Mundo!")

}
func formHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "POST request successful")
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {

	fmt.Printf("Starting server at port 8080\n")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/messages", helloWorldFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
