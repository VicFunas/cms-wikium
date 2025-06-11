package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Monster struct {
	ID          primitive.ObjectID   `bson:"_id"`
	Name        string               `bson:"name"`
	LocationIDs []primitive.ObjectID `bson:"locationIDs"`
}
