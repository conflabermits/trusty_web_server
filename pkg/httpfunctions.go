package httpfunctions

import (
	"fmt"
	"net/http"
	"os"
)

func loadText(file string) string {
	content, err := os.ReadFile(file)
	if err != nil {
		os.Exit(7)
	}
	return string(content)
}

func Respond_ok(w http.ResponseWriter, req *http.Request) {
	ok_text := loadText("pkg/appname-OK.json")
	fmt.Fprintf(w, ok_text)
}

func Respond_degraded(w http.ResponseWriter, req *http.Request) {
	degraded_text := loadText("pkg/appname-DEGRADED.json")
	fmt.Fprintf(w, degraded_text)
}

func Respond_outage(w http.ResponseWriter, req *http.Request) {
	outage_text := loadText("pkg/appname-OUTAGE.json")
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
