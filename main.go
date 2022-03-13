package main

import (
	"aletheiaware.com/netgo"
	"aletheiaware.com/netgo/handler"
	"log"
	"net/http"
)

func main() {
	// Setup Logging to File
	logFile, err := netgo.SetupLogging()
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.Println("Log File:", logFile.Name())

	// Create Multiplexer
	mux := http.NewServeMux()

	// Handle Index
	mux.Handle("/", handler.Log(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Hello World!"))
	})))

	// Serve HTTP Requests
	log.Println("HTTP Server Listening on :80")
	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatal(err)
	}
}
