package repository

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type CategoryMongo struct {
	db *mongo.Collection
}

func NewCategoryMongo(db *mongo.Database) *CategoryMongo {
	return &CategoryMongo{
		db: db.Collection("category"),
	}
}

func (r *CategoryMongo) CreateCategory(ctx context.Context, ctg domain.Category) error {
	_, err := r.db.InsertOne(ctx, ctg)
	return err
}

func (r *CategoryMongo) GetAllCategory(ctx context.Context) ([]domain.Category, error) {
	var categoryList []domain.Category

	cur, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		return []domain.Category{}, err
	}

	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var result domain.Category

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}
		categoryList = append(categoryList, result)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return categoryList, nil
}