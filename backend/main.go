// backend/main.go

package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"task-management/backend/api"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Connecting to MongoDB
	mongoURI := "mongodb://mongo:27017"
	// mongoURI := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(context.TODO())

	collection := client.Database("taskdb").Collection("tasks")

	router := mux.NewRouter()
	api.SetupRoutes(router, collection)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	})

	httpHandler := cors.Handler(router)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, httpHandler))
}
