package createlabel

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
)

type execute struct {
	repository repository.IMongoTrayRepo
}

func NewCreateLabel(repo repository.IMongoTrayRepo) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(data *entity.Tray) (int64, error) {
	result, err := e.repository.CreateLabel(data)

	if err != nil {
		return result, err
	}

	return result, err
}
