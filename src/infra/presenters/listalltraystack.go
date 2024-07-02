package presenters

import (
	"github.com/grrlopes/go-looptask/src/domain/entity"
)

func ListAllLabelTrayStackSuccess(label []entity.LabelStack) output {
	return output{
		"error":   nil,
		"message": label,
		"success": true,
	}
}

func ListAllLabelTrayStackError(err error) output {
	return output{
		"error":   err.Error(),
		"message": nil,
		"success": false,
	}
}
