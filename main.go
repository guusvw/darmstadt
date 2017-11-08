package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	once           = &sync.Once{}
	prometheusAddr = flag.String("prometheus-address", ":8085", "The Address on which the prometheus handler should be exposed")
	prometheusPath = flag.String("prometheus-path", "/metrics", "The path on the host, on which the handler is available")
	address        = flag.String("address", ":8080", "The address to listen on")
)

// ServeForever the prometheus metrics endpoint
func ServeForever(addr, path string) {
	once.Do(func() {
		log.Printf("Prometheus metrics exposed on: %s%s", addr, path)

		http.Handle(path, promhttp.Handler())
		log.Fatal(http.ListenAndServe(addr, nil))
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handlerFail(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Hi there, I'm an error!", http.StatusBadRequest)
}

func handlerSlow(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	fmt.Fprint(w, "Hi there, I respondin very slow.")
}

func main() {
	flag.Parse()

	go ServeForever(*prometheusAddr, *prometheusPath)

	log.Printf("Listening on %s", *address)
	http.HandleFunc("/slow/", handlerSlow)
	http.HandleFunc("/fail/", handlerFail)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
