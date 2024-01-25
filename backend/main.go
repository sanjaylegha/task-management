// main.go

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Task model
type Task struct {
	ID          string `json:"id,omitempty" bson:"_id,omitempty"`
	TaskName    string `json:"taskName,omitempty" bson:"taskName,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`
	Status      string `json:"status,omitempty" bson:"status,omitempty"`
}

// Initialize MongoDB connection
func init() {
	mongoURI := "mongodb://mongo:27017"
	// mongoURI := "mongodb://localhost:27017"

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, _ = mongo.Connect(context.TODO(), clientOptions)
}

// API Handlers
func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task Task
	json.NewDecoder(r.Body).Decode(&task)

	// Perform validation, if needed

	// Insert task into MongoDB
	collection := client.Database("taskdb").Collection("tasks")
	resp, err := collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Fatal(err)
	}

	task.ID = resp.InsertedID.(primitive.ObjectID).Hex()

	log.Printf("Inserted task with ID: %s\n", task.ID)

	json.NewEncoder(w).Encode(task)
}

func listTasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task

	collection := client.Database("taskdb").Collection("tasks")
	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	log.Printf("Found %v tasks\n", len(tasks))

	json.NewEncoder(w).Encode(tasks)
}

func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from request params
	taskID := r.URL.Query().Get("id")

	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex ERROR:", err)
	}

	// Delete task from MongoDB
	collection := client.Database("taskdb").Collection("tasks")
	resp, err := collection.DeleteOne(context.TODO(), bson.M{"_id": taskIDObj})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Deleted %v documents\n", resp.DeletedCount)

	if resp.DeletedCount == 0 {
		fmt.Fprintf(w, "Task with ID %s not found", taskID)
		return
	}

	fmt.Fprintf(w, "Task with ID %s deleted successfully", taskID)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var updatedTask Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	// Extract task ID from request params
	taskID := r.URL.Query().Get("id")
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex ERROR:", err)
	}

	// Update task in MongoDB
	collection := client.Database("taskdb").Collection("tasks")
	filter := bson.M{"_id": taskIDObj}
	update := bson.D{{"$set", updatedTask}}
	resp, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated %v documents\n", resp.ModifiedCount)

	updatedTask.ID = taskID

	json.NewEncoder(w).Encode(updatedTask)
}

func updateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from request params
	taskID := r.URL.Query().Get("id")

	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex ERROR:", err)
	}

	// Extract new status from request body
	var statusUpdate struct {
		Status string `json:"status"`
	}
	json.NewDecoder(r.Body).Decode(&statusUpdate)

	// Update status in MongoDB
	collection := client.Database("taskdb").Collection("tasks")
	filter := bson.D{{"_id", taskIDObj}}
	update := bson.D{{"$set", bson.D{{"status", statusUpdate.Status}}}}
	resp, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if resp.ModifiedCount == 0 {
		log.Printf("Task with ID %s not found", taskID)
	}

	fmt.Fprintf(w, "Status of task with ID %s updated to %s", taskID, statusUpdate.Status)
}

func main() {
	// Define API routes
	http.HandleFunc("/api/tasks/create", createTaskHandler)
	http.HandleFunc("/api/tasks/list", listTasksHandler)
	http.HandleFunc("/api/tasks/delete", deleteTaskHandler)
	http.HandleFunc("/api/tasks/update", updateTaskHandler)
	http.HandleFunc("/api/tasks/update-status", updateTaskStatusHandler)

	// Start HTTP server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server is running on :%s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
