package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"rapigo/internal/db"
	"rapigo/internal/routes"
)

func main() {
	//connect mongodb
	if err := db.Init(); err != nil {
		log.Fatal("Error initializing MongoDB:", err)
	}
	defer db.GetClient().Disconnect(context.TODO())

	routes.SetupRoutes()

	// start the server
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("Server Start Port Number : 8080")
	server.ListenAndServe()
}
