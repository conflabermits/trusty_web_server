package httpfunctions

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

//go:embed static
var content embed.FS

func loadText(file string) string {

	contents, err := content.ReadFile(file)
	if err != nil {
		os.Exit(7)
	}
	return string(contents)
}

func DelayMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		delay := r.URL.Query().Get("delay")
		if delay != "" {
			delaySeconds, err := strconv.Atoi(delay)
			if err == nil {
				time.Sleep(time.Duration(delaySeconds) * time.Second)
			}
		}
		next.ServeHTTP(w, r)
	})
}

func Respond_ok(w http.ResponseWriter, req *http.Request) {
	ok_text := loadText("static/appname-OK.json")
	fmt.Fprintln(w, ok_text)
}

func Respond_degraded(w http.ResponseWriter, req *http.Request) {
	degraded_text := loadText("static/appname-DEGRADED.json")
	fmt.Fprintln(w, degraded_text)
}

func Respond_outage(w http.ResponseWriter, req *http.Request) {
	outage_text := loadText("static/appname-OUTAGE.json")
	fmt.Fprintln(w, outage_text)
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
	http.Redirect(w, req, "/200", http.StatusMovedPermanently)
}

func Respond_302(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "/ok", http.StatusFound)
}

func Respond_headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func Respond_delay(w http.ResponseWriter, req *http.Request) {
	delay := req.URL.Query().Get("delay")
	if delay == "" {
		delay = "1"
	}
	delaySeconds, err := strconv.Atoi(delay)
	if err != nil {
		http.Error(w, "Invalid delay value", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Delaying for %d seconds starting at %s\n", delaySeconds, time.Now().Format(time.RFC3339))
	time.Sleep(time.Duration(delaySeconds) * time.Second)
	fmt.Fprintf(w, "Done delaying at %s\n", time.Now().Format(time.RFC3339))
}
