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
// one collections - TrayAggSet
type LabelAggSet struct {
	ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Trays      []TrayAggSet       `json:"trays" bson:"trays"`
	TrayCount  int64              `json:"tray_count" bson:"tray_count"`
	SmallCount int32              `json:"small_count" bson:"small_count"`
	LargeCount int32              `json:"large_count" bson:"large_count"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

// LabelStackAggSet represents an aggregation result-set for
// one collection - CreatorAggSet
type LabelStackAggSet struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Owner    CreatorAggSet      `json:"owner" bson:"owner"`
	Estimate struct {
		Small int32 `json:"small" bson:"small"`
		Large int32 `json:"large" bson:"large"`
	} `json:"estimate"`
	TrayCount  int64     `json:"tray_count" bson:"tray_count"`
	SmallCount int32     `json:"small_count" bson:"small_count"`
	LargeCount int32     `json:"large_count" bson:"large_count"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time `json:"updated_at" bson:"updated_at"`
}
