package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Labeled struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Trays     []Tray             `json:"trays" bson:"trays"`
	Owner     primitive.ObjectID `json:"owner" bson:"owner"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type Count struct {
	Total_rows int64 `json:"total_rows" validate:"numeric" bson:"total_rows"`
	Offset     int64 `json:"offset" validate:"numeric" bson:"offset"`
}

type MongoResul struct {
	ID        interface{} `json:"id" bson:"_id"`
	TotalRows int         `json:"total_rows"`
	Offset    int         `json:"offset"`
	Rows      []Tray      `json:"rows"`
	Error     string      `json:"error"`
	Reason    string      `json:"reason"`
}

// LabelAggSet represents an aggregation result-set for
// two collections - TrayAggSet, CreatorAggSet
type LabelAggSet struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Trays     []TrayAggSet       `json:"trays" bson:"trays"`
	Owner     CreatorAggSet      `json:"owner" bson:"owner"`
	TrayCount int64              `json:"tray_count" bson:"tray_count"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

// LabelStack represents an struct without Trays slice element
type LabelStack struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Owner     primitive.ObjectID `json:"owner" bson:"owner"`
	TrayCount int64              `json:"tray_count" bson:"tray_count"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
