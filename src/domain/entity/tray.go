package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tray struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	TrayId string             `json:"trayid" bson:"trayid"`
	Size   string             `json:"size" bson:"size"`
	UserId primitive.ObjectID `json:"userid" bson:"userid"`
	Done   bool               `json:"done" bson:"done"`
}
