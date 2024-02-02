package main

import (
    "net/http"
    "lesson/httpadapter"
    "lesson/databaseadapter"
)

func main() {
	// Instantiate core components
	userRepo := &databaseadapter.DatabaseUserAdapter{}

	// Instantiate adapters and wire them to the core
	httpAdapter := &httpadapter.HttpUserAdapter{CorePort: userRepo}

	// Set up HTTP routes
	http.HandleFunc("/user", httpAdapter.GetUser)

	// Start the server
	http.ListenAndServe(":8080", nil)
}