package listalltraystack

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
)

type execute struct {
	repository repository.IMongoTrayRepo
}

func NewListAllTrayStack(repo repository.IMongoTrayRepo) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute() ([]entity.LabelStackAggSet, error) {
	result, err := e.repository.ListAllTrayStack()
	if err != nil {
		return []entity.LabelStackAggSet{}, err
	}

	return result, nil
}
