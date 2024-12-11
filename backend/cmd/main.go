package main

import (
	"band-manager-backend/internal/db"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("BACKEND_PORT")
	if port == "" {
		port = "8080"
	}

	db.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// database := db.GetDB()
		fmt.Fprintf(w, "Hello World!")
	})

	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}