package repository

import "github.com/grrlopes/go-looptask/src/domain/entity"

type IMongoUserRepo interface {
	FindUserByName(data *entity.Users) (entity.Users, error)
	CreateUser(data *entity.Users) (entity.MongoResul, error)
	FindUserByEmailandUser(data *entity.Users) (entity.Users, error)
	FindUserById(data *entity.Users) (entity.Users, error)
}
