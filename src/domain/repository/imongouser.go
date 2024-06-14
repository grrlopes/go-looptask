package repository

import "github.com/grrlopes/go-looptask/src/domain/entity"

type IMongoUserRepo interface {
	FindUserByName(data *entity.Users) (entity.Users, error)
	UserSave(data *entity.Users) (entity.MongoResul, error)
}
