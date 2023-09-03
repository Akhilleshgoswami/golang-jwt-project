package main

import (
	"fmt"
	"net/http"
	routes "todo-api/internal/api/handlers"
)

func main() {
	http.HandleFunc("/hello", routes.helloHandler)

	port := ":8000"

	fmt.Printf("Server is listenin on port %s ...\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
