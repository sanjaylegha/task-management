package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"task-management/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ApiHandler struct {
	collection *mongo.Collection
}

func NewApiHandler(collection *mongo.Collection) *ApiHandler {
	return &ApiHandler{collection: collection}
}

func (ah *ApiHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating Tasks ...")
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	// Perform validation, if needed

	// Insert task into MongoDB

	resp, err := ah.collection.InsertOne(context.TODO(), task)
	if err != nil {
		log.Fatal(err)
	}

	task.ID = resp.InsertedID.(primitive.ObjectID)

	log.Printf("Inserted task with ID: %s\n", task.ID)

	json.NewEncoder(w).Encode(task)
}

func (ah *ApiHandler) ListTasksHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Listing tasks ...")
	var tasks []models.Task

	cursor, err := ah.collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		tasks = append(tasks, task)
	}

	log.Printf("Found %v tasks\n", len(tasks))

	json.NewEncoder(w).Encode(tasks)
}

func (ah *ApiHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting Tasks ...")
	// Extract task ID from request params
	taskID := r.URL.Query().Get("id")

	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex ERROR:", err)
	}

	// Delete task from MongoDB

	resp, err := ah.collection.DeleteOne(context.TODO(), bson.M{"_id": taskIDObj})
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

func (ah *ApiHandler) UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating Tasks ...")

	var updatedTask models.Task
	json.NewDecoder(r.Body).Decode(&updatedTask)

	// Extract task ID from request params
	taskID := r.URL.Query().Get("id")
	taskIDObj, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		fmt.Println("primitive.ObjectIDFromHex ERROR:", err)
	}

	// Update task in MongoDB

	filter := bson.M{"_id": taskIDObj}
	update := bson.D{{"$set", updatedTask}}
	resp, err := ah.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Updated %v documents\n", resp.ModifiedCount)

	updatedTask.ID = taskIDObj

	json.NewEncoder(w).Encode(updatedTask)
}

func (ah *ApiHandler) UpdateTaskStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating Task status ...")
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

	filter := bson.D{{"_id", taskIDObj}}
	update := bson.D{{"$set", bson.D{{"status", statusUpdate.Status}}}}
	resp, err := ah.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if resp.ModifiedCount == 0 {
		log.Printf("Task with ID %s not found", taskID)
	}

	fmt.Fprintf(w, "Status of task with ID %s updated to %s", taskID, statusUpdate.Status)
}
