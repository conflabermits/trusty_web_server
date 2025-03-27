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
		log.Printf("Request received for %s\n", r.URL)
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
					log.Printf("Fail check: FAIL if failratePercent > randomNumber. Requested failratePercent is %d. System generated randomNumber %d.\n", failratePercent, randomNumber)
					if failratePercent > randomNumber {
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

func RegisterStatusCodeHandlers(mux *http.ServeMux) {
	for code := 100; code < 600; code++ {
		mux.Handle(fmt.Sprintf("/%d", code), http.HandlerFunc(func(code int) http.HandlerFunc {
			return func(w http.ResponseWriter, r *http.Request) {
				//log.Printf("Processing for code %d\n", code)
				switch code {
				case 218:
					w.WriteHeader(218)
					fmt.Fprintf(w, "This is fine\nhttps://en.wikipedia.org/wiki/List_of_HTTP_status_codes#Unofficial_codes\n")
				case 301, 302, 303, 307, 308:
					http.Redirect(w, r, "/200", code)
				case 420:
					w.WriteHeader(420)
					fmt.Fprintf(w, "Enhance your calm\nhttps://en.wikipedia.org/wiki/List_of_HTTP_status_codes#Unofficial_codes\n")
				case 530:
					w.WriteHeader(530)
					fmt.Fprintf(w, "Site is frozen\nhttps://en.wikipedia.org/wiki/List_of_HTTP_status_codes#Unofficial_codes\n")
				default:
					http.Error(w, http.StatusText(code), code)
				}
			}
		}(code)))
	}
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

func Respond_headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
