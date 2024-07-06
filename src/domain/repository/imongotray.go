package repository

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

type IMongoTrayRepo interface {
	CreateLabelTray(data *entity.Labeled) (string, error)
	CreateLabel(data *entity.Tray) (int64, error)
	ListAllTrays(data *entity.Labeled) (entity.MongoResul, error)
	Fetchtraybyid(data *entity.TrayId) ([]entity.LabelAggSet, error)
	ListAllTrayStack() ([]entity.LabelStack, error)
}
