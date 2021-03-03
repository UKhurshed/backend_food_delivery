package repository

import (
	"backend_food_delivery/pkg/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Collection
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db.Collection("user")}
}

func (r *AuthMongo) CreateUser(ctx context.Context, user domain.User) error {
	_, err := r.db.InsertOne(ctx, user)
	return err

}

func (r *AuthMongo) GetUser(ctx context.Context, email, password string) (domain.User, error) {
	var user domain.User
	if err := r.db.FindOne(ctx, bson.M{"email": email, "password": password}).Decode(&user); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.User{}, err
		}
		return domain.User{}, err
	}
	return user, nil
}
