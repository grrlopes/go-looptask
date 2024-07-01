package fetchtraybyid

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.TrayId) (entity.LabelAggSet, error)
}
