package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Author    string             `json:"author" validate:"required,min=4,max=100" bson:"author"`
	Email     string             `json:"email" validate:"required,email" bson:"email"`
	Password  string             `json:"password" validate:"required" bson:"password"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	jwt.StandardClaims
}
