package mongodb

import (
	"context"
	"errors"
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
			{Key: "$addFields", Value: bson.D{
				{Key: "tray_count", Value: bson.D{
					{Key: "$size", Value: "$trays"},
				}},
			}},
		},
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
func (db *trays) CreateLabelTray(data *entity.LabelTrayStack) (string, error) {
  fmt.Println(data.Large, data.Small)
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
			Key: "estimate", Value: bson.D{
				{Key: "small", Value: data.Small},
				{Key: "large", Value: data.Large},
			},
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

func (db *trays) ListAllTrayStack() ([]entity.LabelStackAggSet, error) {
	var result []entity.LabelStackAggSet

	pipeline := bson.A{
		// Lookup the 'owner' details from the 'user' collection
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
		// Unwind the 'owner' array (in case it's an array) and preserve non-null results only
		bson.D{
			{Key: "$unwind",
				Value: bson.D{
					{Key: "path", Value: "$owner"},
					{Key: "preserveNullAndEmptyArrays", Value: false},
				},
			},
		},
		// Add a new field 'tray_count' to count the number of elements in the 'trays' array
		bson.D{
			{Key: "$addFields", Value: bson.D{
				{Key: "tray_count", Value: bson.D{
					{Key: "$size", Value: "$trays"},
				}},
			}},
		},
		// Facet the pipeline into two paths: one for documents with trays and one for those without
		bson.D{{Key: "$facet", Value: bson.D{
			// Pipeline for documents with non-empty trays
			{Key: "with_trays", Value: bson.A{
				// Unwind the 'trays' array to process each tray individually
				bson.D{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$trays"}}}},
				// Group by document _id to keep original structure and count tray sizes
				bson.D{{Key: "$group", Value: bson.D{
					{Key: "_id", Value: "$_id"},
					{Key: "small_count", Value: bson.D{{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$trays.size", "small"}}}, 1, 0}}}}}},
					{Key: "large_count", Value: bson.D{{Key: "$sum", Value: bson.D{
						{Key: "$cond", Value: bson.A{bson.D{{Key: "$eq", Value: bson.A{"$trays.size", "large"}}}, 1, 0}}}}}},
					{Key: "document", Value: bson.D{{Key: "$first", Value: "$$ROOT"}}},
				}}},
				// Add the counts to the original document structure
				bson.D{{Key: "$addFields", Value: bson.D{
					{Key: "document.small_count", Value: "$small_count"},
					{Key: "document.large_count", Value: "$large_count"},
				}}},
				// Replace the root with the modified document
				bson.D{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$document"}}}},
			}},
			// Pipeline for documents without trays
			{Key: "without_trays", Value: bson.A{
				// Match documents where the trays array is empty
				bson.D{{Key: "$match", Value: bson.D{{Key: "tray_count", Value: 0}}}},
			}},
		}}},
		// Merge the results from both facets
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "result", Value: bson.D{{Key: "$concatArrays", Value: bson.A{"$with_trays", "$without_trays"}}}},
		}}},
		bson.D{{Key: "$unwind", Value: "$result"}},
		bson.D{{Key: "$replaceRoot", Value: bson.D{{Key: "newRoot", Value: "$result"}}}},
		// Sort the documents by the 'created_at' field
		bson.D{{Key: "$sort", Value: bson.D{{Key: "created_at", Value: -1}}}}, // 1 for ascending order, -1 for descending order
		// Project the desired fields to include in the final result
		bson.D{
			{Key: "$project", Value: bson.D{
				{Key: "trays", Value: 0},
				{Key: "owner.updated_at", Value: 0},
				{Key: "owner.created_at", Value: 0},
				{Key: "owner.password", Value: 0},
			}},
		},
		bson.D{{Key: "$skip", Value: 0}},
		bson.D{{Key: "$limit", Value: 14}},
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

func (db *trays) CreateLabel(data *entity.Tray) (int64, error) {
	labeled := entity.Tray{
		Id:        primitive.NewObjectID(),
		TrayId:    data.TrayId,
		Size:      data.Size,
		UserId:    data.UserId,
		Done:      data.Done,
		CreatedAt: time.Now(),
	}

	pipeline := bson.M{
		"$push": bson.M{
			"trays": labeled,
		},
	}

	result, err := db.con.UpdateOne(context.TODO(), bson.M{"_id": data.Id}, pipeline)
	if err != nil {
		return 0, errors.New(err.Error())
	}

	return result.MatchedCount, nil
}
