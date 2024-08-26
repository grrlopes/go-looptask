package fetchtraystackbydate

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute() ([]entity.LabelStackAggSet, error)
}
