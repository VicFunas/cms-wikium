package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type PerfectLink struct {
	ModID     primitive.ObjectID `bson:"_id"`
	MonsterID primitive.ObjectID `bson:"monsterID"`
}

type Character struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	PerfectLinks []PerfectLink      `bson:"perfectLinks"`
}
