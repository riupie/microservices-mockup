package main

import (
	"fmt"
	"log"
	"net/http"
  "os"
)

func viewIndex(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("content-type", "text/plain; charset=utf-8")
	w.Write([]byte(`Hello from Golang Backend`))
}

func viewPing(w http.ResponseWriter, r *http.Request) {
  hostname, err := os.Hostname()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  fmt.Fprintf(w, "pong from %s", hostname)
}

func router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/home", viewIndex)
	mux.HandleFunc("/ping", viewPing)
	return mux
}

func main() {
	port := 8080
	log.Printf("Running on http://127.0.0.1:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router()))
}
