package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID    primitive.ObjectID `bson: "_id , omitempty"`
	Title string             `bson  : "title"`
	Done  bool               `bson:"done"`
}