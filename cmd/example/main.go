package main

import (
	"flag"
	"net/http"
	"embed"
	"fmt"
	"log"
)

//go:embed index.html example.*
var FS embed.FS

func main() {

	host := flag.String("host", "localhost", "The host name to listen for requests on")
	port := flag.Int("port", 8080, "The host port to listen for requests on")
	
	flag.Parse()

	mux := http.NewServeMux()

	http_fs := http.FS(FS)
	example_handler := http.FileServer(http_fs)

	mux.Handle("/", example_handler)
	
	addr := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("Listening for requests on %s\n", addr)

	err := http.ListenAndServe(addr, mux)

	if err != nil {
		log.Fatalf("Failed to start server, %v", err)
	}
}
