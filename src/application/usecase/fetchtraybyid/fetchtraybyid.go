package fetchtraybyid

import (
	"errors"

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

func (e execute) Execute(data *entity.TrayId) (entity.LabelAggSet, error) {
	var res entity.LabelAggSet
	result, err := e.repository.Fetchtraybyid(data)

	if len(result) <= 0 {
		return entity.LabelAggSet{}, errors.New("Not found record!")
	}

	if err != nil {
		return res, errors.New("Something went wrong!")
	}

	for _, v := range result {
		res = v
	}

	return res, err
}
