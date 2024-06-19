package createlabeltray

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Labeled) (entity.MongoResul, error)
}
