package fetchuserbyid

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Users) (entity.Users, error)
}
