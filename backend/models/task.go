// models/task.go
package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Task represents a task model
type Task struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	TaskName    string             `json:"taskName" bson:"taskName"`
	Description string             `json:"description" bson:"description"`
	Status      string             `json:"status" bson:"status"`
}
