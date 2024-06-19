package mongodb

import (
	"context"

	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
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
func (db *trays) ListAllTrays(data *entity.Labeled) (entity.MongoResul, error) {
	panic("unimplemented")
}
func (db *trays) CreateLabelTray(data *entity.Labeled) (entity.MongoResul, error) {
	pipeline := bson.D{
		{
			Key: "id", Value: data.ID,
		},
		{
			Key: "tray", Value: bson.D{
				{
					Key: "id", Value: data.Trays.Id,
				},
				{
					Key: "trayid", Value: data.Trays.TrayId,
				},
				{
					Key: "size", Value: data.Trays.Size,
				},
				{
					Key: "done", Value: data.Trays.Done,
				},
			},
		},
	}

	_, err := db.con.InsertOne(context.TODO(), pipeline)
	if err != nil {
		return entity.MongoResul{}, err
	}

	var result entity.MongoResul

	return result, err
}
