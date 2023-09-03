package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"todo-api/routes"
	"todo-api/utils"
)

func main() {
	// Initialize MongoDB connection
	err := utils.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to initialize MongoDB: %v", err)
	}

	// Define the port for your Go web server
	port := "8080"

	router := routes.SetupRouter()
	// Start your HTTP server
	server := &http.Server{
		Addr:         ":" + port,
		Handler: router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("Server is listening on port %s...\n", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error: %v\n", err)
	}
}
