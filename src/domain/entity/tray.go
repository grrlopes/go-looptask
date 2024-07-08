package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tray struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	TrayId    string             `json:"trayid" bson:"trayid"`
	Size      string             `json:"size" bson:"size"`
	UserId    primitive.ObjectID `json:"userid" bson:"userid"`
	Done      bool               `json:"done" bson:"done"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

type LabelId struct {
	Id string `form:"id"`
}

type TrayId struct {
	Id primitive.ObjectID `json:"id" validate:"required" bson:"_id"`
}

// TrayAggSet represents an aggregation result-set for one collection Creator
type TrayAggSet struct {
	Id        primitive.ObjectID `json:"id" bson:"_id"`
	TrayId    string             `json:"trayid" bson:"trayid"`
	Size      string             `json:"size" bson:"size"`
	UserId    CreatorAggSet      `json:"userid" bson:"userid"`
	Done      bool               `json:"done" bson:"done"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
