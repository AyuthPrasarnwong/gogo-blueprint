package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewConnection(clientOptions *options.ClientOptions) (*mongo.Client, error) {
	db, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(context.Background(), readpref.Primary()); err != nil {
		return nil, err
	}

	return db, nil
}
