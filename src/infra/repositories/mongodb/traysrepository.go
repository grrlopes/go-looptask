package mongodb

import (
	"context"
	"errors"
	"time"

	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type trays struct {
	con *mongo.Collection
}

func NewTrayRepository() repository.IMongoTrayRepo {
	err := OpenDB()
	if err != nil {
		panic(err)
	}

	db := GetDBCollection("tray")

	return &trays{
		con: db,
	}
}

func (db *trays) Fetchtraybyid(data *entity.Tray) (entity.Labeled, error) {
	var result entity.Labeled
	err := db.con.FindOne(context.TODO(), bson.D{{
		Key:   "_id",
		Value: data.Id,
	}}).Decode(&result)

	if len(result.Trays) == 0 {
		return result, errors.New("No record found")
	}

	if err != nil {
		return result, err
	}

	return result, nil
}

func (db *trays) ListAllTrays(data *entity.Labeled) (entity.MongoResul, error) {
	panic("unimplemented")
}
func (db *trays) CreateLabelTray(data *entity.Labeled) (string, error) {
	trays := bson.A{}
	for _, tray := range data.Trays {
		trays = append(trays, bson.D{
			{Key: "_id", Value: primitive.NewObjectID()},
			{Key: "trayid", Value: tray.TrayId},
			{Key: "size", Value: tray.Size},
			{Key: "userid", Value: tray.UserId},
			{Key: "done", Value: tray.Done},
		})
	}
	pipeline := bson.D{
		{
			Key: "trays", Value: trays,
		},
		{
			Key: "owner", Value: data.Owner,
		},
		{
			Key: "created_at", Value: time.Now(),
		},
		{
			Key: "updated_at", Value: time.Now(),
		},
	}

	res, err := db.con.InsertOne(context.TODO(), pipeline)
	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), err
}
