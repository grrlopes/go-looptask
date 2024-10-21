package entity

type ValidateJwt struct {
	Token string `json:"token" validate:"omitempty,required" bson:"token"`
}
