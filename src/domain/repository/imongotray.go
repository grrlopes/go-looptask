package repository

import "github.com/grrlopes/go-looptask/src/domain/entity"

type IMongoTrayRepo interface {
	ListAllTrays(data *entity.Tray) (entity.MongoResul, error)
}
