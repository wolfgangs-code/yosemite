package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// ANSI color codes
	const (
		RESET = "\033[0m"
		GREEN = "\033[32m"
		BLUE  = "\033[34m"
		GREY  = "\033[90m"
	)
	fmt.Println(GREEN + "Yosemite" + RESET + " Backend Server" + BLUE + " (Valley)" + GREY + " v0.0.0" + RESET)

	// Serve Port
	port := os.Getenv("YOSEMITE_PORT")
	if port == "" {
		port = "8080"
	}

	// Create router
	r := mux.NewRouter()
	defineRoutes(r)

	// Start the server
	fmt.Printf("Listening on port "+GREEN+":%s"+RESET+"...\n", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
