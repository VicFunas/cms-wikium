package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Location struct {
	ID      primitive.ObjectID `bson:"_id"`
	Name    string             `bson:"name"`
	Kingdom []Kingdom          `bson:"kingdoms"`
}
