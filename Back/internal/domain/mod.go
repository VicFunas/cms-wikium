package domain

import "go.mongodb.org/mongo-driver/v2/bson"

type Mod struct {
	ID   bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name string        `bson:"name"`
}
