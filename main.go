package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	err := LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	err = InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer CloseDB()

	r := gin.Default()
	SetupRoutes(r)

	err = r.Run()
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}
