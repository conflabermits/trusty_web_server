package main

import (
	"flag"
	"log"
	"net/http"

	"httpfunctions"
)

func main() {
	port := flag.String("port", "8080", "Port to run the local web server")
	flag.Parse()

	mux := http.NewServeMux()
	httpfunctions.RegisterStatusCodeHandlers(mux)

	mux.Handle("/ok", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_ok)))
	mux.Handle("/degraded", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_degraded)))
	mux.Handle("/outage", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_outage)))
	mux.Handle("/headers", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_headers)))

	log.Printf("Server starting on http://localhost:%s\n", *port)
	if err := http.ListenAndServe(":"+*port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
