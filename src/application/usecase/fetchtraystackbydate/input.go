package fetchtraystackbydate

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute(startDate *entity.TrayStacked) ([]entity.LabelStackAggSet, error)
}
