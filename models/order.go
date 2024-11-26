package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	UserID      primitive.ObjectID `bson:"user_id"`
	Origin      string             `bson:"origin"`
	Destination string             `bson:"destination"`
	Distance    int                `bson:"distance"`
	Price       float64            `bson:"price"`
	Status      string             `bson:"status"`
	CreatedAt   primitive.DateTime `bson:"created_at"`
}
