package main

import (
	"net/http"
	
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, _ = w.Write([]byte("This is a test"))
	})
	_ = http.ListenAndServe(":6060", http.DefaultServeMux)
}
