
package listalltrays

import "github.com/grrlopes/go-looptask/src/domain/entity"

type InputBoundary interface {
	Execute(data *entity.Tray) (entity.MongoResul, error)
}
