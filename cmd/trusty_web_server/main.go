package main

import (
    "flag"
    "net/http"

    "github.com/trusty_web_server/httpfunctions"
)

func main() {

    port := flag.String("port", "8080", "Port to run the local web server")
    flag.Parse()

    http.HandleFunc("/ok", httpfunctions.respond_ok)
    http.HandleFunc("/degraded", httpfunctions.respond_degraded)
    http.HandleFunc("/outage", httpfunctions.respond_outage)

    http.HandleFunc("/200", httpfunctions.respond_200)
    http.HandleFunc("/301", httpfunctions.respond_301)
    http.HandleFunc("/302", httpfunctions.respond_302)
    http.HandleFunc("/401", httpfunctions.respond_401)
    http.HandleFunc("/404", httpfunctions.respond_404)
    http.HandleFunc("/500", httpfunctions.respond_500)

    http.HandleFunc("/headers", httpfunctions.headers)

    http.ListenAndServe(":"+*port, nil)
}
