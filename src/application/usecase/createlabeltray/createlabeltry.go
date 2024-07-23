package createlabeltray

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
)

type execute struct {
	repository repository.IMongoTrayRepo
}

func NewListAllTrays(repo repository.IMongoTrayRepo) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute(data *entity.LabelTrayStack) (string, error) {
	result, err := e.repository.CreateLabelTray(data)

	if err != nil {
		return result, err
	}

	return result, err
}
