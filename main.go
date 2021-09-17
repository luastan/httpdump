package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

// Loggers setup
func init() {
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	ErrorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

// Prints request to stdout and sends the response
func requestHandler(w http.ResponseWriter, req *http.Request) {
	InfoLogger.Printf("Request from %s:", req.RemoteAddr)
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		w.WriteHeader(404)
		ErrorLogger.Println("Error dumping the request")
		return
	} else {
		println(string(requestDump))
		_, err := fmt.Fprintf(w, "<html><title>httpdump</title><body><pre>\n%s\n</pre></body></html>\n", string(requestDump))
		if err != nil {
			ErrorLogger.Println("Error writing the response")
		}
	}

}

func main() {
	// Argument parsing
	flag.Parse()
	listenAddr := "0.0.0.0:8080"
	if flag.NArg() > 0 {
		listenAddr = flag.Arg(0)
	}

	// Listener set-up
	http.HandleFunc("/", requestHandler)
	InfoLogger.Printf("Listening on %s ...\n", listenAddr)
	err := http.ListenAndServe(listenAddr, nil)

	// Error handling
	if err != nil {
		ErrorLogger.Fatalf("Could not set-up listener on %s", listenAddr)
	}

}
