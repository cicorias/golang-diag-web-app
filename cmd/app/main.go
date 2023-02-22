package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func emitHeaders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Request Headers\n")
	for name, headers := range r.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "\t%v: %v\n", name, h)
		}
	}

	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "\tRequest Details\n")
	fmt.Fprintf(w, "\tMethod:   %v\n", r.Method)
	fmt.Fprintf(w, "\tHost:     %v\n", r.Host)
	fmt.Fprintf(w, "\tURL:      %v\n", r.URL)
	fmt.Fprintf(w, "\tURI:      %v\n", r.RequestURI)
	fmt.Fprintf(w, "\tProto:    %v\n", r.Proto)
	fmt.Fprintf(w, "\tRemote:   %v\n", r.RemoteAddr)
	fmt.Fprintf(w, "\tUsername: %v\n", r.URL.User.Username())
	fmt.Fprintf(w, "\tLength:   %v\n", r.ContentLength)

	fmt.Fprintf(w, "Env Details\n")
	fmt.Fprintf(w, "\tHOSTNAME: %v\n", hostname)

}

func emitEnvironment(w http.ResponseWriter, r *http.Request) {
	env := os.Environ()
	for _, e := range env {
		fmt.Fprintf(w, "%v\n", e)
	}

}

func main() {
	http.HandleFunc("/other", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/", emitHeaders)
	http.HandleFunc("/env", emitEnvironment)

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
