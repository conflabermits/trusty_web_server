package httpfunctions

import (
	"fmt"
	"net/http"
)

func Respond_ok(w http.ResponseWriter, req *http.Request) {
	ok_text := `{
    "components": [
        {
            "description": "Most important check",
            "essential": true,
            "name": "auth-service",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:38080/auth-service/health"
        },
        {
            "description": "Less important check",
            "essential": false,
            "name": "activity-webservice",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:38080/activity-service/health"
        },
        {
            "description": "Some other cheeck",
            "essential": true,
            "name": "database",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:48080/user-table"
         }
    ],
    "name": "appname",
    "statusCode": "OK"
}`
	fmt.Fprintf(w, ok_text)
}

func Respond_degraded(w http.ResponseWriter, req *http.Request) {
	degraded_text := `{
    "components": [
        {
            "description": "Most important check",
            "essential": true,
            "name": "auth-service",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:38080/auth-service/health"
        },
        {
            "description": "Less important check",
            "essential": false,
            "name": "activity-webservice",
            "statusCode": "CRITICAL",
            "statusText": "Can't reach activity service, returns 404",
            "uri": "http://localhost:38080/activity-service/health"
        },
        {
            "description": "Some other cheeck",
            "essential": true,
            "name": "database",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:48080/user-table"
        }
    ],
    "name": "appname",
    "statusCode": "DEGRADED"
}`
	fmt.Fprintf(w, degraded_text)
}

func Respond_outage(w http.ResponseWriter, req *http.Request) {
	outage_text := `{
    "components": [
        {
            "description": "Most important check",
            "essential": true,
            "name": "auth-service",
            "statusCode": "CRITICAL",
            "statusText": "Can't reach auth service, returns 500",
            "uri": "http://localhost:38080/auth-service/health"
        },
        {
            "description": "Less important check",
            "essential": false,
            "name": "activity-webservice",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:38080/activity-service/health"
        },
        {
            "description": "Some other cheeck",
            "essential": true,
            "name": "database",
            "statusCode": "OK",
            "statusText": null,
            "uri": "http://localhost:48080/user-table"
        }
    ],
    "name": "appname",
    "statusCode": "OUTAGE"
}`
	fmt.Fprintf(w, outage_text)
}

func Respond_200(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "\n")
}

func Respond_401(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(401)
	fmt.Fprintf(w, "Not authorized\n")
}

func Respond_404(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(404)
	fmt.Fprintf(w, "Page not found\n")
}

func Respond_500(w http.ResponseWriter, req *http.Request) {

	w.WriteHeader(500)
	fmt.Fprintf(w, "Internal server error\n")
}

func Respond_301(w http.ResponseWriter, req *http.Request) {

	http.Redirect(w, req, "/200", 301)

}

func Respond_302(w http.ResponseWriter, req *http.Request) {

	http.Redirect(w, req, "/ok", 302)

}

func Respond_headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
