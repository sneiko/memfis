package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"memfis/internal/handlers"
)

func main() {
	var filename string
	flag.StringVar(&filename, "file", "", "Path to memory profile file")
	flag.Parse()

	handler := handlers.NewHandler()

	if filename == "" {
		log.Fatalf("No file provided. Visit http://localhost:8080 to upload a memory profile file\n")
	}

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		log.Fatalf("File does not exist: %s", filename)
	}

	if err := handler.LoadFile(filename); err != nil {
		log.Fatalf("Error loading file: %v", err)
	}

	fmt.Printf("Loaded memory profile: %s\n", filename)

	// Setup routes
	http.HandleFunc("/", handler.IndexHandler)

	// API endpoints
	http.HandleFunc("/api/data", handler.APIDataHandler)
	http.HandleFunc("/api/memstats", handler.APIMemStatsHandler)
	http.HandleFunc("/api/stacktraces", handler.APIStackTracesHandler)

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Memory Leak Analyzer started on http://0.0.0.0:8080")
	if filename != "" {
		fmt.Println("Analyzing file:", filename)
	} else {
		fmt.Println("Visit http://localhost:8080 to upload a memory profile file")
	}

	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
