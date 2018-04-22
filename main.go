package main

import (
    "flag"
    "log"
    "net/http"

    "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
    
    var port string
	flag.StringVar(&port, "port", "8000", "port number")
	flag.Parse()
	
    http.Handle("/metrics", promhttp.Handler())
    http.Handle("/", http.FileServer(http.Dir("/var/www/html/public")))
    // log.Fatal(http.ListenAndServe(":8000", nil))
    http.ListenAndServe(":"+port, nil)
}
