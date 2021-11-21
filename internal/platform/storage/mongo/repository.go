package mongo

import (
	"context"
	"time"

	freegames "github.com/arkiant/freegames/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type repository struct {
	client     *mongo.Client
	database   string
	collection string
	timeout    time.Duration
}

func newMongoClient(mongoURL string, mongoTimeout time.Duration) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), mongoTimeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewMongoRepository create a new mongo repository
func NewMongoRepository(mongoURL, mongoDB string, mongoTimeout time.Duration) (freegames.GameRepository, error) {
	repo := &repository{
		timeout:    mongoTimeout,
		database:   mongoDB,
		collection: "currentFreeGames",
	}
	client, err := newMongoClient(mongoURL, mongoTimeout)
	if err != nil {
		return nil, err
	}
	repo.client = client
	return repo, nil
}

// GetGames get all current free games
func (r *repository) GetGames(platform freegames.Platform) (freegames.FreeGames, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	fg := make(freegames.FreeGames, 0)

	collection := r.client.Database(r.database).Collection(r.collection)
	cur, err := collection.Find(ctx, bson.D{{"platform", platform.GetName()}})
	if err != nil {
		return fg, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		game := freegames.Game{}
		err := cur.Decode(&game)
		if err != nil {
			return fg, err
		}
		fg = append(fg, game)
	}

	return fg, nil
}

// Exists check if a game exists in database
func (r *repository) Exists(game freegames.Game) bool {

	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	collection := r.client.Database(r.database).Collection(r.collection)
	filter := bson.M{"name": game.Name, "platform": game.Platform}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return false
	}

	return count > 0
}

// Store a free game into the database
func (r *repository) Store(game freegames.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.InsertOne(
		ctx,
		game,
	)
	if err != nil {
		return err
	}
	return nil
}

// Delete a old free game from the database
func (r *repository) Delete(game freegames.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()
	collection := r.client.Database(r.database).Collection(r.collection)

	_, err := collection.DeleteOne(ctx, bson.D{{"name", game.Name}})
	if err != nil {
		return err
	}
	return nil
}
