package utils

import (
	"context"
	"log"
	"time"
	"todo-api/models"

	"go.mongodb.org/mongo-driver/bson"
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
	var todos []models.Todo
	for cursor.Next(ctx) {

		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)

	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return todos, nil
}
