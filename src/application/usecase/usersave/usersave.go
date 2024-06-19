package usersave

import (
	"errors"
	"fmt"
	"time"

	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
	"github.com/grrlopes/go-looptask/src/helper"
)

type execute struct {
	findRepository repository.IMongoUserRepo
}

func NewUserSave(repo repository.IMongoUserRepo) InputBoundary {
	return execute{
		findRepository: repo,
	}
}

func (e execute) Execute(data *entity.Users) (entity.MongoResul, error) {
	var resul entity.MongoResul
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.Password, _ = helper.CreatePassword(data)

	res, _ := e.findRepository.FindUserByEmailandUser(data)

	if res.Email != "" {
		error := errors.New(fmt.Sprintf("%s", "This email is already used."))
		resul.Error = error.Error()
		return resul, error
	}

	result, err := e.findRepository.UserSave(data)

	if err != nil {
		return entity.MongoResul{}, err
	}

	if result.Error == "unauthorized" {
		error := errors.New(result.Error)
		return result, error
	}

	return result, nil
}
