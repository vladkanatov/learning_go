package main

import (
	"log"
	"todo-app/pkg/api"
	"todo-app/pkg/storage"
)

func main() {
	db := storage.SetupDatabase()
	router := api.SetupRouter(db)

	log.Println("Starting server on :8000")
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
