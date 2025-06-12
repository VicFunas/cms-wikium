package repository

import (
	"context"

	"github.com/VicFunas/cms-wikium/internal/domain"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ModRepository struct {
	collection *mongo.Collection
}

func NewModRepository(db *mongo.Database) *ModRepository {
	return &ModRepository{
		collection: db.Collection("mods"),
	}
}

// GetModByID finds a mod by their ID string.
func (r *ModRepository) GetModByID(ctx context.Context, id string) (domain.Mod, error) {
	var mod domain.Mod

	// Convert the ID string from the URL to a MongoDB ObjectID.
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return domain.Mod{}, err // Invalid ID format
	}

	// Find the document in the collection.
	filter := bson.M{"_id": objectID}
	err = r.collection.FindOne(ctx, filter).Decode(&mod)
	if err != nil {
		return domain.Mod{}, err // Handles "no documents in result" (not found)
	}

	return mod, nil
}

// CreateMod inserts a new mod into the database.
func (r *ModRepository) CreateMod(ctx context.Context, mod domain.Mod) (domain.Mod, error) {
	// The driver will generate a new ObjectID if mod.ID is zero.
	result, err := r.collection.InsertOne(ctx, mod)
	if err != nil {
		return domain.Mod{}, err
	}

	// Set the generated ID on the mod struct to return to the caller.
	mod.ID = result.InsertedID.(bson.ObjectID)
	return mod, nil
}
