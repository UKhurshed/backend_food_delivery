package repository

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type TodoListMongo struct {
	db *mongo.Collection
}

func NewTodoListMongo(db *mongo.Database) *TodoListMongo {
	return &TodoListMongo{db: db.Collection("todo")}
}

func (r *TodoListMongo) Create(ctx context.Context, list domain.Todo) error {
	_, err := r.db.InsertOne(ctx, list)
	return err
}

func (r *TodoListMongo) GetAll(ctx context.Context) ([]domain.Todo, error) {
	var lists []domain.Todo

	cur, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		return []domain.Todo{}, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result domain.Todo

		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		lists = append(lists, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return lists, nil
}
