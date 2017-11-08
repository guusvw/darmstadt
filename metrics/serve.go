package metrics

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Serve the prometheus metrics endpoint
func Serve(addr, path string) {
	log.Printf("Metrics exposed on %s%s", addr, path)

	http.Handle(path, promhttp.Handler())
	log.Fatal(http.ListenAndServe(addr, nil))
}
