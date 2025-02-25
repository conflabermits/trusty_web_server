# Trusty Web Server

This is a simple web server with a simple goal: to respond with a consistent set of responses (HTTP response code, content, and/or actions) based on the endpoint requested.

It can be used for testing programs that make requests to a web server. The value is that the responses are consistent and predictable, which allows you to validate the functionality of the program making the requests instead of second-guessing if the remote endpoint is responding reliably.

## Requests and Expected Responses

| Request | HTTP Response Code | Response Content | Action |
| --- | --- | --- | --- |
| [/ok](http://localhost:8080/ok) | 200 | Health check response: "OK" |  |
| [/degraded](http://localhost:8080/degraded) | 200 | Health check response: "Degraded" |  |
| [/outage](http://localhost:8080/outage) | 200 | Health check response: "Outage" |  |
| [/200](http://localhost:8080/200) | 200 | (Blank page) |  |
| [/301](http://localhost:8080/301) | 301 | | Redirect to "/200" |
| [/302](http://localhost:8080/302) | 302 | | Redirect to "/ok" |
| [/401](http://localhost:8080/401) | 401 | "Not authorized" |  |
| [/404](http://localhost:8080/404) | 404 | "Page not found" |  |
| [/500](http://localhost:8080/500) | 500 | "Internal server error" |  |
| [/headers](http://localhost:8080/headers) | 200 | Client headers |  |
| [?delay=N](http://localhost:8080/200?delay=3) | * | * | Delay for N seconds |
| [?failrate=N](http://localhost:8080/200?failrate=50) | * | * | N% of requests will fail |

## Running the server

The web server can be run as-is using `go run`:

```text
$ go run main.go
2025/02/24 22:18:30 Server starting on http://localhost:8080
```

It can also be compiled into a binary for easier portability:

```text
$ go build && ls -ln trusty_web_server
-rwxr-xr-x 1 1000 1000 6579433 Oct  9 16:39 trusty_web_server
```

In either case it will run on port 8080 by default.

It takes `-port` as an optional argument, allowing you to specify which port the web server should be served on:

```text
$ go run main.go -port 48080
2025/02/24 22:18:30 Server starting on http://localhost:48080
```

```text
$ ./trusty_web_server -port 48080
2025/02/24 22:18:30 Server starting on http://localhost:48080
```

## Resources and References

* Stream Notes [002](https://github.com/conflabermits/Scripts/blob/main/stream/pilot/002/notes.md), [013](https://github.com/conflabermits/Scripts/blob/main/stream/pilot/013/notes.md)
* [Example HTTP server code](https://gobyexample.com/http-servers)
* [List of HTTP status codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes)
* [How to set HTTP status code in response in Go](https://golangbyexample.com/set-http-status-code-golang/)
* [Example of setting HTTP redirect in Golang](https://gist.github.com/hSATAC/5343225)
