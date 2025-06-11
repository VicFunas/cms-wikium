package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Story struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
