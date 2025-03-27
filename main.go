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

	mux.Handle("/ok", http.HandlerFunc(httpfunctions.Respond_ok))
	mux.Handle("/degraded", http.HandlerFunc(httpfunctions.Respond_degraded))
	mux.Handle("/outage", http.HandlerFunc(httpfunctions.Respond_outage))
	mux.Handle("/headers", http.HandlerFunc(httpfunctions.Respond_headers))

	handler := httpfunctions.HTTPMiddleware(mux)

	log.Printf("Server starting on http://localhost:%s\n", *port)
	if err := http.ListenAndServe(":"+*port, handler); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
