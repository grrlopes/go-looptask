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

func (e execute) Execute(data *entity.Labeled) (entity.MongoResul, error) {
	result, _ := e.repository.CreateLabelTray(data)

	return result, nil
}
