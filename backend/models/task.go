// models/task.go
package models

import "gopkg.in/mgo.v2/bson"

// Task represents a task model
type Task struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	TaskName    string        `json:"taskName" bson:"taskName"`
	Description string        `json:"description" bson:"description"`
	Status      string        `json:"status" bson:"status"`
}
