package api

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes configures API routes
func SetupRoutes(router *mux.Router, collection *mongo.Collection) {
	apiHandler := NewApiHandler(collection)

	// Add middleware
	// router.Use(CorsMiddleware)

	// Define API routes
	router.HandleFunc("/api/tasks/create", apiHandler.CreateTaskHandler).Methods("POST")
	router.HandleFunc("/api/tasks/list", apiHandler.ListTasksHandler).Methods("GET")
	router.HandleFunc("/api/tasks/delete", apiHandler.DeleteTaskHandler).Methods("DELETE")
	router.HandleFunc("/api/tasks/update", apiHandler.UpdateTaskHandler).Methods("PUT")
	router.HandleFunc("/api/tasks/update-status", apiHandler.UpdateTaskStatusHandler).Methods("PUT")
}
