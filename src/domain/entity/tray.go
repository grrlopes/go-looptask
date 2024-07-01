package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Tray struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	TrayId string             `json:"trayid" bson:"trayid"`
	Size   string             `json:"size" bson:"size"`
	UserId primitive.ObjectID `json:"userid" bson:"userid"`
	Done   bool               `json:"done" bson:"done"`
}

type TrayId struct {
	ID primitive.ObjectID `json:"id" validate:"required" bson:"_id"`
}

// TrayAggSet represents an aggregation result-set for one collection Creator
type TrayAggSet struct {
	Id     primitive.ObjectID `json:"id" bson:"_id"`
	TrayId string             `json:"trayid" bson:"trayid"`
	Size   string             `json:"size" bson:"size"`
	UserId CreatorAggSet          `json:"userid" bson:"userid"`
	Done   bool               `json:"done" bson:"done"`
}
