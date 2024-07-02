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

func (db *trays) Fetchtraybyid(data *entity.TrayId) ([]entity.LabelAggSet, error) {
	var result []entity.LabelAggSet

	pipeline := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "_id", Value: data.Id}}}},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "user"},
					{Key: "localField", Value: "owner"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "owner"},
				},
			},
		},
		bson.D{
			{Key: "$unwind",
				Value: bson.D{
					{Key: "path", Value: "$owner"},
					{Key: "preserveNullAndEmptyArrays", Value: false},
				},
			},
		},
		bson.D{
			{Key: "$lookup",
				Value: bson.D{
					{Key: "from", Value: "user"},
					{Key: "localField", Value: "trays.userid"},
					{Key: "foreignField", Value: "_id"},
					{Key: "as", Value: "tray_user"},
				},
			},
		},
		bson.D{
			{Key: "$unwind", Value: bson.D{
				{Key: "path", Value: "$tray_user"},
				{Key: "preserveNullAndEmptyArrays", Value: false},
			}},
		},
		bson.D{
			{Key: "$addFields", Value: bson.D{
				{Key: "trays.userid", Value: "$tray_user"},
			}}},
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "tray_user", Value: 0},
				{Key: "trays.userid.password", Value: 0},
				{Key: "owner.updated_at", Value: 0},
				{Key: "owner.created_at", Value: 0},
				{Key: "owner.password", Value: 0},
				{Key: "trays.userid.updated_at", Value: 0},
				{Key: "trays.userid.created_at", Value: 0},
			}},
		},
	}

	cursor, err := db.con.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return result, err
	}

	defer cursor.Close(context.TODO())

	err = cursor.All(context.TODO(), &result)
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

func (db *trays) ListAllTrayStack() ([]entity.LabelStack, error) {
	var result []entity.LabelStack

	res, err := db.con.Find(context.TODO(), bson.D{})
	if err != nil {
		return result, errors.New(err.Error())
	}

	defer res.Close(context.TODO())

	err = res.All(context.TODO(), &result)
	if err != nil {
		return result, errors.New(err.Error())
	}

	return result, nil
}
