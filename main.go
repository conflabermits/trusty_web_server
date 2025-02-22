package main

import (
	"flag"
	"fmt"
	"net/http"

	"httpfunctions"
)

func main() {

	port := flag.String("port", "8080", "Port to run the local web server")
	flag.Parse()

	http.HandleFunc("/ok", httpfunctions.Respond_ok)
	http.HandleFunc("/degraded", httpfunctions.Respond_degraded)
	http.HandleFunc("/outage", httpfunctions.Respond_outage)

	http.HandleFunc("/200", httpfunctions.Respond_200)
	http.HandleFunc("/301", httpfunctions.Respond_301)
	http.HandleFunc("/302", httpfunctions.Respond_302)
	http.HandleFunc("/401", httpfunctions.Respond_401)
	http.HandleFunc("/404", httpfunctions.Respond_404)
	http.HandleFunc("/500", httpfunctions.Respond_500)

	http.HandleFunc("/headers", httpfunctions.Respond_headers)

	fmt.Printf("Server starting on http://localhost:%s\n", *port)
	http.ListenAndServe(":"+*port, nil)
}
