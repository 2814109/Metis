package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello World!!! test")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("0.0.0.0:18888", nil))
}
