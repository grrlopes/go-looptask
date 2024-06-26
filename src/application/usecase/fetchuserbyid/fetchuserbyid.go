package fetchuserbyid

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
)

type execute struct {
	repository repository.IMongoUserRepo
}

func NewFetchOneUser(repo repository.IMongoUserRepo) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(data *entity.Users) (entity.Users, error) {
	result, err := e.repository.FindUserById(data)

	return result, err
}
