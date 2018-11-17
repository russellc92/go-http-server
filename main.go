package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)
func hostname() string {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	return name
}

func handler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "GO HTTP Server, Running on host: " + hostname())
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
