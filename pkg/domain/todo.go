package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Price       string             `json:"price" binding:"required"`
	Title       string             `json:"title" binding:"required"`
	Description string             `json:"description" binding:"required"`
	//Available   bool               `json:"available" binding:"required"`
}
