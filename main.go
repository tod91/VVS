package main

import (
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("hello there and welcome"))
}

func main() {
	http.HandleFunc("/", mainPage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
