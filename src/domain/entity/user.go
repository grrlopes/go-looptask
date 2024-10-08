package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Name      string             `json:"name" validate:"required,min=4,max=100" bson:"name"`
	Surname   string             `json:"surname" validate:"required,min=4,max=100" bson:"surname"`
	Email     string             `json:"email" validate:"required,email" bson:"email"`
	Password  string             `json:"password" validate:"required" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	jwt.StandardClaims
}

type UserId struct {
	ID primitive.ObjectID `json:"id" validate:"omitempty,required" bson:"_id"`
}

// CreatorAggSet represents an aggregation result-set
type CreatorAggSet struct {
	ID      primitive.ObjectID `json:"id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	Surname string             `json:"surname" bson:"surname"`
	Email   string             `json:"email" bson:"email"`
}
