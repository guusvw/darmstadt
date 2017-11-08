package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/guusvw/darmstadt/metrics"
)

var (
	once           = &sync.Once{}
	prometheusAddr = flag.String("prometheus-address", ":8085", "The Address on which the prometheus handler should be exposed")
	prometheusPath = flag.String("prometheus-path", "/metrics", "The path on the host, on which the handler is available")
	address        = flag.String("address", ":8080", "The address to listen on")

	// now provides func() time.Time
	// so it is easier to mock, if wou want to add tests
	now = time.Now
)

func handler(w http.ResponseWriter, r *http.Request) {
	start := now()

	metrics.IncConnections()

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

	metrics.ConnectionTime(now().Sub(start))
}

func handlerFail(w http.ResponseWriter, r *http.Request) {
	start := now()

	metrics.IncErrors(metrics.Fatal)

	metrics.IncConnections()

	http.Error(w, "Hi there, I'm an error!", http.StatusBadRequest)

	metrics.ConnectionTime(now().Sub(start))
}

func handlerInfo(w http.ResponseWriter, r *http.Request) {
	start := now()

	metrics.IncErrors(metrics.Info)

	metrics.IncConnections()

	http.Error(w, "Hi there, I'm an error!", http.StatusBadRequest)

	metrics.ConnectionTime(now().Sub(start))
}

func handlerSlow(w http.ResponseWriter, r *http.Request) {
	start := now()

	metrics.IncConnections()

	time.Sleep(3 * time.Millisecond)
	fmt.Fprint(w, "Hi there, I respondin very slow.")

	metrics.ConnectionTime(now().Sub(start))
}

func main() {
	flag.Parse()

	go metrics.Serve(*prometheusAddr, *prometheusPath)

	log.Printf("Listening on %s", *address)
	http.HandleFunc("/slow/", handlerSlow)
	http.HandleFunc("/fail/", handlerFail)
	http.HandleFunc("/info/", handlerInfo)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
