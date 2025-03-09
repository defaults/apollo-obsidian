package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "build/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
