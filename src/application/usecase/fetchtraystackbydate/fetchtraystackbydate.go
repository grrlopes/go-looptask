package fetchtraystackbydate

import (
	"time"

	"github.com/grrlopes/go-looptask/src/domain/entity"
	"github.com/grrlopes/go-looptask/src/domain/repository"
)

type execute struct {
	repository repository.IMongoTrayRepo
}

func NewFetchtrayStackByDate(repo repository.IMongoTrayRepo) InputBoundary {
	return execute{
		repository: repo,
	}
}

func (e execute) Execute() ([]entity.LabelStackAggSet, error) {
	var (
		now       = time.Now()
		startDate = time.Date(
			now.Year(),
			now.Month(),
			10,
			0, 0, 0, 0,
			time.UTC,
		)
		endDate = startDate.AddDate(0, 0, 60)
	)

	result, err := e.repository.FetchTrayStackByDate(startDate, endDate)
	if err != nil {
		return []entity.LabelStackAggSet{}, err
	}

	return result, nil
}
