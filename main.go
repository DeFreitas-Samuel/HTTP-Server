package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", helloWorldFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func helloWorldFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Received")
	fmt.Fprintf(w, "Hello World Hola Mundo! '%v'", r.Method)

}
