package mongodb

import (
	"context"
	"fmt"

	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type users struct {
	con *mongo.Collection
}

func NewUserRepository() repository.IMongoUserRepo {
	err := OpenDB()
	if err != nil {
		panic(err)
	}

	db := GetDBCollection("user")

	return &users{
		con: db,
	}
}

func (db *users) UserSave(data *entity.Users) (entity.MongoResul, error) {
	pipeline := bson.D{
		{
			Key: "author", Value: data.Author,
		},
		{
			Key: "email", Value: data.Email,
		},
		{
			Key: "password", Value: data.Password,
		},
		{
			Key: "created_at", Value: data.CreatedAt,
		},
		{
			Key: "updated_at", Value: data.UpdatedAt,
		},
	}
	fmt.Println(data.Email)

	_, err := db.con.InsertOne(context.TODO(), pipeline)
	fmt.Println(err, "ppppp##{{{{}}}}")
	if err != nil {
		return entity.MongoResul{}, err
	}

	var result entity.MongoResul
	result.Reason = "created!"

	return result, err
}

func (db *users) FindUserByName(user *entity.Users) (entity.Users, error) {
	var result entity.Users
	err := db.con.FindOne(context.TODO(), bson.D{{
		Key:   "author",
		Value: user.Author,
	}}).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}
