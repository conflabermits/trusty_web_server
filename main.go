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

	http.Handle("/ok", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_ok)))
	http.Handle("/degraded", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_degraded)))
	http.Handle("/outage", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_outage)))

	http.Handle("/200", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_200)))
	http.Handle("/301", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_301)))
	http.Handle("/302", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_302)))
	http.Handle("/401", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_401)))
	http.Handle("/404", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_404)))
	http.Handle("/500", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_500)))

	http.Handle("/headers", httpfunctions.DelayMiddleware(http.HandlerFunc(httpfunctions.Respond_headers)))
	http.HandleFunc("/delay", httpfunctions.Respond_delay)

	fmt.Printf("Server starting on http://localhost:%s\n", *port)
	http.ListenAndServe(":"+*port, nil)
}
