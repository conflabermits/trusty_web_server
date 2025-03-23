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

	http.Handle("/ok", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_ok)))
	http.Handle("/degraded", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_degraded)))
	http.Handle("/outage", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_outage)))
	/*
		http.Handle("/200", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_200)))
		http.Handle("/301", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_301)))
		http.Handle("/302", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_302)))
		http.Handle("/401", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_401)))
		http.Handle("/404", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_404)))
		http.Handle("/500", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_500)))
	*/
	http.Handle("/headers", httpfunctions.HTTPMiddleware(http.HandlerFunc(httpfunctions.Respond_headers)))

	mux := http.NewServeMux()
	httpfunctions.RegisterStatusCodeHandlers(mux)

	log.Printf("Server starting on http://localhost:%s\n", *port)
	if err := http.ListenAndServe(":"+*port, mux); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	//http.ListenAndServe(":"+*port, nil)
}
