package fetchtraystackbydate

import (
	"errors"
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

func (e execute) Execute(dateStart *entity.TrayStacked) ([]entity.LabelStackAggSet, error) {
	var (
		now       = dateStart.CreatedAt
		startDate = time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			0, 0, 0, 0,
			time.UTC,
		)
		endDate = startDate.AddDate(0, 0, 60)
	)

	result, err := e.repository.FetchTrayStackByDate(startDate, endDate)
  if len(result) == 0 {
		return []entity.LabelStackAggSet{}, errors.New("There is no record on this date")
  }
	if err != nil {
		return []entity.LabelStackAggSet{}, err
	}

	return result, nil
}
