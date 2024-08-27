package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LabelTrayStack struct {
	Small int32              `json:"small" validate:"required,min=1,max=300" bson:"small"`
	Large int32              `json:"large" validate:"required,min=1,max=300" bson:"large"`
	Owner primitive.ObjectID `json:"owner" bson:"owner"`
	Trays []Tray             `json:"trays" bson:"trays"`
}

type TrayStacked struct {
	CreatedAt time.Time `json:"created_at" validate:"required,notZeroTime" bson:"created_at"`
}
