package createlabeltray

import (
	"errors"
	"time"

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
	var (
		now       = time.Now()
		startDate = time.Date(
			now.Year(),
			now.Month(),
			now.Day(),
			0, 0, 0, 0,
			time.UTC,
		)
		endDate = startDate.AddDate(0, 0, 1)
	)

	isDuplicated, _ := e.repository.FetchTrayStackByDate(startDate, endDate)
	if len(isDuplicated) > 0 {
		return "", errors.New("Only one stack tray is allowed per day")
	}

	result, err := e.repository.CreateLabelTray(data)

	if err != nil {
		return result, err
	}

	return result, err
}
