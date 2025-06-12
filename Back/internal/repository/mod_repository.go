package repository

import (
	"context"
	"errors"

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

// GetAllMods finds all mods in database
func (r *ModRepository) GetAllMods(ctx context.Context) ([]domain.Mod, error) {
	var mods []domain.Mod

	// Find the document in the collection.
	// err = r.collection.FindOne(ctx, filter).Decode(&mod)
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return []domain.Mod{}, err // Handles "no documents in result" (not found)
	}
	defer cursor.Close(ctx)

	err = cursor.All(ctx, &mods)
	if err != nil {
		return []domain.Mod{}, err // Handles "no documents in result" (not found)
	}

	return mods, nil
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

// UpdateMod updates an existing mod in the database.
func (r *ModRepository) UpdateMod(ctx context.Context, id string, mod domain.Mod) (domain.Mod, error) {
	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return domain.Mod{}, err
	}

	filter := bson.M{"_id": objectID}

	update := bson.M{
		"$set": bson.M{
			"name":        mod.Name,
			"description": mod.Description,
		},
	}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return domain.Mod{}, err
	}

	if result.MatchedCount == 0 {
		return domain.Mod{}, errors.New("mod not found")
	}

	return mod, nil
}
