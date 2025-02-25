package httpfunctions

import (
	"embed"
	"fmt"
	"log"
	"math/rand"
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

func HTTPMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		delay := r.URL.Query().Get("delay")
		if delay != "" {
			delaySeconds, err := strconv.Atoi(delay)
			if err == nil {
				log.Printf("Delaying request for %d seconds.\n", delaySeconds)
				w.Header().Set("X-Delay-Seconds", strconv.Itoa(delaySeconds))
				time.Sleep(time.Duration(delaySeconds) * time.Second)
			}
		}
		failrate := r.URL.Query().Get("failrate")
		if failrate != "" {
			failratePercent, err := strconv.Atoi(failrate)
			if err == nil {
				if failratePercent > 0 && failratePercent <= 100 {
					randomNumber := rand.Intn(100) + 1
					w.Header().Set("X-Failrate-Percent", strconv.Itoa(failratePercent))
					w.Header().Set("X-Random-Number", strconv.Itoa(randomNumber))
					log.Printf("Fail check. Request must beat a %d to fail. Request rolled a D100 and got %d.\n", randomNumber, failratePercent)
					if randomNumber < failratePercent {
						w.WriteHeader(500)
						fmt.Fprintf(w, "Internal server error - Request failed successfully\n")
						return
					}
				}
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
