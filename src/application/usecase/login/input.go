package login

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Users) (OutputBoundary, error)
}
