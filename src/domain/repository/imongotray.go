package repository

import "github.com/grrlopes/go-looptask/src/domain/entity"

type IMongoTrayRepo interface {
	CreateLabelTray(data *entity.Labeled) (entity.MongoResul, error)
	ListAllTrays(data *entity.Labeled) (entity.MongoResul, error)
}
