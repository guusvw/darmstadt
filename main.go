package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	address = flag.String("address", ":8080", "The address to listen on")
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	flag.Parse()

	log.Printf("Listening on %s", *address)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
