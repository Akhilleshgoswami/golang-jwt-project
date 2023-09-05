package utils

import (
	"context"
	"fmt"
	"log"
	"time"
	"todo-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	dbName     = "todoapp"
	collection = "todos"
)

func InitMongoDB() error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clintOption := options.Client().ApplyURI("mongodb+srv://Akhil132:PpECv88WJ4Ck8Xgn@cluster0.bxrd3dz.mongodb.net/?retryWrites=true&w=majority")
	var err error
	client, err = mongo.Connect(ctx, clintOption)

	if err != nil {
		return err

	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	log.Println("Connected to MongoDb")
	return nil
}

func GetDB() *mongo.Database {
	return client.Database(dbName)
}
func GetTodosCollection() *mongo.Collection {
	return GetDB().Collection(collection)
}

func AddTodoToDB(todo *models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := GetTodosCollection().InsertOne(ctx, todo)

	if err != nil {
		return err

	}
	return nil
}

func GetAllTodosFromDB() ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	cursor, err := GetTodosCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	fmt.Println(cursor)
	var todos []models.Todo
	for cursor.Next(ctx) {

		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		fmt.Println(todo)
		todos = append(todos, todo)

	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todos, nil

}

// UpdateTodoInDB updates a todo item in the MongoDB database.
func UpdateTodoInDB(todoID primitive.ObjectID, updatedTodo *models.Todo) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define the filter to find the todo item by its ID
	filter := bson.M{"_id": todoID}
	fmt.Println("data id ", todoID)

	// Define the update to set the new values
	update := bson.M{
		"$set": bson.M{
			"title": updatedTodo.Title,
			"done":  updatedTodo.Done,
		},
	}

	// Configure the update options
	updateOptions := options.Update().SetUpsert(false)

	// Update the todo item in the database
	result, err := GetTodosCollection().UpdateOne(ctx, filter, update, updateOptions)
	if err != nil {
		return err
	}

	// Check if the update affected any documents
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments // Return a custom error if no documents were matched
	}

	return nil
}

// DeleteTodoFromDB deletes a todo item from the MongoDB database by its ID.
func DeleteTodoFromDB(todoID primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Define the filter to find the todo item by its ID
	filter := bson.M{"_id": todoID}

	// Delete the todo item from the database
	result, err := GetTodosCollection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	// Check if the delete operation affected any documents
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments // Return a custom error if no documents were deleted
	}

	fmt.Println("data deleted")
	return nil
}
