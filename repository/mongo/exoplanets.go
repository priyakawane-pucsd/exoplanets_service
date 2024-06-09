package mongo

import (
	"context"
	"errors"
	"exoplanetservice/logger"
	"exoplanetservice/models/dao"
	"exoplanetservice/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

const (
	EXOPLANETS_COLLECTION = "exoplanets"
)

func (r *Repository) CreateExoplanets(ctx context.Context, req *dao.Exoplanets) error {
	collection := r.conn.Database(r.cfg.Database).Collection(EXOPLANETS_COLLECTION)
	// Convert created_at and updated_at to milliseconds
	req.CreatedAt = time.Now().UnixNano() / int64(time.Millisecond)
	req.UpdatedAt = req.CreatedAt //Assume created_at and updated_at are the same initially
	// Set the ID of the question
	objectID := primitive.NewObjectID().Hex()
	req.ID = objectID
	// Insert the exoplanet document into the collection
	_, err := collection.InsertOne(ctx, req)
	if err != nil {
		logger.Error(ctx, "Error inserting exoplanet: %v", err)
		return utils.NewInternalServerError("Failed to insert exoplanet into the database")
	}
	return nil
}

// GetExoplanets retrieves a list of exoplanets from the database.
func (r *Repository) GetExoplanets(ctx context.Context, limit, offset int) ([]*dao.Exoplanets, error) {
	collection := r.conn.Database(r.cfg.Database).Collection(EXOPLANETS_COLLECTION)
	findOptions := options.Find().SetLimit(int64(limit)).SetSkip(int64(offset))
	cursor, err := collection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, utils.NewInternalServerError("Failed to retrieve exoplanets from the database")
	}
	defer cursor.Close(ctx)

	var exoplanets []*dao.Exoplanets
	if err := cursor.All(ctx, &exoplanets); err != nil {
		return nil, utils.NewInternalServerError("Failed to decode exoplanets")
	}
	return exoplanets, nil
}

func (r *Repository) GetExoplanetById(ctx context.Context, id string) (*dao.Exoplanets, error) {
	collection := r.conn.Database(r.cfg.Database).Collection(EXOPLANETS_COLLECTION)
	filter := bson.M{"_id": id}
	var result *dao.Exoplanets
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, utils.NewBadRequestError("Exoplanet not found")
		}
		logger.Error(ctx, "Error finding exoplanet by ID: %v", err)
		return nil, utils.NewInternalServerError("Failed to retrieve exoplanet by ID")
	}
	return result, nil
}

func (r *Repository) UpdateExoplanetById(ctx context.Context, exoplanet *dao.Exoplanets, exoplanetId string) error {
	collection := r.conn.Database(r.cfg.Database).Collection(EXOPLANETS_COLLECTION)
	filter := bson.M{"_id": exoplanetId}
	update := bson.M{"$set": exoplanet}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error(ctx, "Error updating exoplanet by ID: %v", err)
		return utils.NewInternalServerError("Failed to update exoplanet")
	}
	return nil
}

func (r *Repository) DeleteExoplanetById(ctx context.Context, id string) error {
	collection := r.conn.Database(r.cfg.Database).Collection(EXOPLANETS_COLLECTION)
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Error(ctx, "Error deleting exoplanet by ID: %v", err)
		return utils.NewInternalServerError("Failed to delete exoplanet")
	}
	// Check if any document was deleted
	if res.DeletedCount == 0 {
		logger.Error(ctx, "No exoplanet found with ID: %s", id)
		return utils.NewBadRequestError("Exoplanet not found")
	}
	return nil
}
