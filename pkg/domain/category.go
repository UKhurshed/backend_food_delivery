package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" binding:"required"`
	Image string `json:"image" binding:"required"`
}

