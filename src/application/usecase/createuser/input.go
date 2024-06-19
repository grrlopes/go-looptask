package createuser

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Users) (entity.MongoResul, error)
}
