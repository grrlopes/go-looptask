package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
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
	var result []entity.LabelUser

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
		// bson.D{
		// 	{Key: "$lookup",
		// 		Value: bson.D{
		// 			{Key: "from", Value: "user"},
		// 			{Key: "localField", Value: "trays.userid"},
		// 			{Key: "foreignField", Value: "_id"},
		// 			{Key: "as", Value: "trays"},
		// 		},
		// 	},
		// },
		// bson.D{
		// 	{Key: "$unwind",
		// 		Value: bson.D{
		// 			{Key: "path", Value: "$trays"},
		// 			{Key: "preserveNullAndEmptyArrays", Value: false},
		// 		},
		// 	},
		// },
	}

	cursor, err := db.con.Aggregate(context.TODO(), pipeline)
	if err = cursor.All(context.TODO(), &result); err != nil {
		panic(err)
	}
	// fmt.Println(result, "+++++++++")
	// for _, result := range result {
	// 	fmt.Printf("Result: %+v\n", result)
	// }
  jk, _ := json.MarshalIndent(result, "", " ")
  fmt.Println(string(jk))

	if err != nil {
		return entity.Labeled{}, err
	}

	return entity.Labeled{}, nil
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
