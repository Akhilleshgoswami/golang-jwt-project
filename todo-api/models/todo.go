package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Todo represents a todo item
type Todo struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title"`
	Done  bool               `json:"done" bson:"done"`
}
