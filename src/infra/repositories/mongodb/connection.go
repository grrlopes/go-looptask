package mongodb

import (
	"context"
	"errors"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Database

func GetDBCollection(coll string) *mongo.Collection {
	return db.Collection(coll)
}

func OpenDB() error {
	uri := os.Getenv("MONGO_URI")

	if uri == "" {
		return errors.New("There is no Mongo env set up!")
	}

	client, err := mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(uri),
	)

	if err != nil {
		return err
	}

	db = client.Database(os.Getenv("SCHEMA"))

	return nil
}

func CloseDB() error {
	con := db.Client()
	return con.Disconnect(context.Background())
}
