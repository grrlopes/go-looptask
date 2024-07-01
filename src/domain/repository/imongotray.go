package repository

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

type IMongoTrayRepo interface {
	CreateLabelTray(data *entity.Labeled) (string, error)
	ListAllTrays(data *entity.Labeled) (entity.MongoResul, error)
	Fetchtraybyid(data *entity.Tray) ([]entity.LabelAggSet, error)
}
