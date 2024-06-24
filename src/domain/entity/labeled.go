package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Labeled struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Trays     []Tray             `json:"trays" bson:"trays"`
	Owner     string             `json:"owner" bson:"owner"`
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

type rows struct {
	ID string `json:"id"`
}
