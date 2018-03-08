package main

import (
	"log"
	"net/http"
)

var dumbHTTP = func(w http.ResponseWriter, r *http.Request) {
	log.Printf("serving %s to %s", r.RequestURI, r.RemoteAddr)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("stupid is what stupid does"))
}

func main() {
	http.HandleFunc("/dumb", dumbHTTP)
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("failed to start http server: %v", err)
	}
}
