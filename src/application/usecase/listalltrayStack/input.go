package listalltraystack

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute() ([]entity.LabelStack, error)
}
