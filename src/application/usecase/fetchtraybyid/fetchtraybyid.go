package fetchtraybyid

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
)

type execute struct {
	repository repository.IMongoTrayRepo
}

func NewFetchOneTray(repo repository.IMongoTrayRepo) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(data *entity.Tray) (entity.Labeled, error) {
	result, err := e.repository.Fetchtraybyid(data)

	return result, err
}
