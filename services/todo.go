package services

import (
	"context"
	"log"
	"time"

	"github.com/codepnw/go-mongo-todos/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type ITodos interface {
	InsertTodo(req *models.Todo) (*models.Todo, error)
	FindAllTodos() ([]*models.Todo, error)
	FindTodoById(id string) (*models.Todo, error)
	UpdateTodo(id string, req *models.Todo) (*mongo.UpdateResult, error)
	DeleteTodo(id string) error
}

type todoService struct {
	collection *mongo.Collection
}

func NewTodoService(collection *mongo.Collection) ITodos {
	return &todoService{collection: collection}
}

func (t *todoService) InsertTodo(req *models.Todo) (*models.Todo, error) {
	todo := &models.Todo{
		Task:      req.Task,
		Completed: false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	result, err := t.collection.InsertOne(context.TODO(), todo)
	if err != nil {
		log.Println("Error:", err)
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	res, err := t.FindTodoById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (t *todoService) FindAllTodos() ([]*models.Todo, error) {
	ctx := context.Background()
	var todos []*models.Todo

	todo, err := t.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer todo.Close(ctx)

	for todo.Next(ctx) {
		var t *models.Todo
		todo.Decode(&t)

		todos = append(todos, t)
	}
	return todos, nil
}

func (t *todoService) FindTodoById(id string) (*models.Todo, error) {
	todo := &models.Todo{}

	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &models.Todo{}, err
	}

	err = t.collection.FindOne(context.Background(), bson.M{"_id": mongoID}).Decode(todo)
	if err != nil {
		return &models.Todo{}, err
	}

	return todo, nil
}

func (s *todoService) UpdateTodo(id string, req *models.Todo) (*mongo.UpdateResult, error) {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.D{
		{Name: "set", Value: bson.D{
			{Name: "task", Value: req.Task},
			{Name: "completed", Value: req.Completed},
			{Name: "updated_at", Value: time.Now()},
		}},
	}

	res, err := s.collection.UpdateOne(context.Background(), bson.M{"_id": mongoID}, update)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

func (s *todoService) DeleteTodo(id string) error {
	mongoID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = s.collection.DeleteOne(context.Background(), bson.M{"_id": mongoID})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
