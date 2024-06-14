package entity

type Tray struct {
	Id     string `json:"id" validate:"required,min=4,max=10" bson:"id"`
	TrayId string `json:"trayid" validate:"required,min=4,max=30" bson:"trayid"`
	Size   string `json:"size" validate:"required,gte=1" bson:"size"`
	User   string `json:"user" validate:"required,gte=1" bson:"user"`
	Done   bool   `json:"done" validate:"required,bool" bson:"done"`
}
