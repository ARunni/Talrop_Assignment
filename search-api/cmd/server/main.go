package main

import (
	"log"
	"net/http"
	"search-api/config"
	"search-api/di"

	_ "github.com/lib/pq"
)

func main() {

	db := config.ConnectDb()
	searchHandler := di.Initialize(db)

	// Define routes
	http.HandleFunc("/search", searchHandler.SearchUsers)

	// Start the server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
